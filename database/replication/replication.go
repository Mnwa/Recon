package replication

import (
	"Recon/backup"
	"Recon/database"
	"github.com/golang/protobuf/proto"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type Replication struct {
	Transmitter chan *Transaction
	Receiver    chan *Transaction

	Received map[string]int64

	Replicas map[string]*fasthttp.HostClient
}

var transactionPool = sync.Pool{
	New: func() interface{} {
		return new(Transaction)
	},
}

func AcquireTransaction() *Transaction {
	return transactionPool.Get().(*Transaction)
}

func ReleaseTransaction(transaction *Transaction) {
	transaction.Reset()
	transactionPool.Put(transaction)
}

func (r *Replication) Receive() {
	for {
		transaction := <-r.Receiver

		go func(transaction *Transaction) {
			for key, data := range transaction.Data {
				if timestamp, ok := r.Received[key]; ok && timestamp <= transaction.Timestamp {
					continue
				}

				database.Client.Set(key, data)

				r.Received[key] = transaction.Timestamp
			}

			ReleaseTransaction(transaction)
		}(transaction)
	}
}

func (r *Replication) Transmit() {
	for {
		transaction := <-r.Transmitter

		if len(r.Replicas) == 0 {
			continue
		}
		go func(transaction *Transaction) {
			body, err := proto.Marshal(transaction)

			if err == nil {
				for addr, host := range r.Replicas {
					req := fasthttp.AcquireRequest()
					req.SetRequestURI("/replication/receiver")
					req.Header.SetMethod("POST")
					req.Header.SetHost(addr)
					req.Header.SetContentType("application/protobuf")
					req.SetBody(body)
					resp := fasthttp.AcquireResponse()

					err := host.Do(req, resp)

					if err != nil {
						log.Println("Master", addr, "is down:", err)
					}

					fasthttp.ReleaseRequest(req)
					fasthttp.ReleaseResponse(resp)
				}
			} else {
				log.Println("Replication error:", err)
			}

			ReleaseTransaction(transaction)
		}(transaction)
	}
}

func (r *Replication) SendMessage(data map[string][]byte) {
	if len(r.Replicas) > 0 {
		r.Transmitter <- NewTransaction(data)
	}
}

func NewReplication(replications []string) *Replication {
	var hosts = make(map[string]*fasthttp.HostClient)
	for _, addr := range replications {
		hosts[addr] = &fasthttp.HostClient{
			Addr: addr,
		}
	}
	return &Replication{
		Transmitter: make(chan *Transaction, 2048),
		Receiver:    make(chan *Transaction, 2048),
		Received:    make(map[string]int64),
		Replicas:    hosts,
	}
}

func NewTransaction(data map[string][]byte) *Transaction {
	transaction := AcquireTransaction()
	transaction.Timestamp = time.Now().UnixNano()
	transaction.Data = data
	return transaction
}

func init() {
	if os.Getenv("RECON_REPLICATION_HOSTS") == "" {
		Replica = NewReplication([]string{})
		return
	}
	replicationHosts := strings.Split(os.Getenv("RECON_REPLICATION_HOSTS"), ",")
	Replica = NewReplication(replicationHosts)
	go Replica.Receive()

	if len(replicationHosts) > 0 {
		go Replica.Transmit()

		if os.Getenv("RECON_REPLICATION_INIT") != "off" {
			for _, host := range Replica.Replicas {
				req := fasthttp.AcquireRequest()
				req.SetRequestURI("/backup")
				req.Header.SetMethod("GET")
				req.Header.SetHost(host.Addr)
				resp := fasthttp.AcquireResponse()

				err := host.Do(req, resp)

				if err == nil {
					err = backup.RestoreBackup(resp.Body())
					fasthttp.ReleaseRequest(req)
					fasthttp.ReleaseResponse(resp)
					if err == nil {
						break
					} else {
						log.Println("Restoring data from master error:", err)
					}
				}
			}
		}
	}
}

var Replica *Replication
