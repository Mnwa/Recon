package replication

import (
	"Recon/backup"
	"Recon/database"
	"github.com/gogo/protobuf/proto"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"strings"
	"time"
)

type Replication struct {
	Transmitter chan Transaction
	Receiver    chan Transaction

	Received map[string]int64

	Replicas map[string]*fasthttp.HostClient
}

func (r *Replication) Receive() {
	var transaction Transaction
	for {
		transaction = <-r.Receiver

		for key, data := range transaction.Data {
			if timestamp, ok := r.Received[key]; ok && timestamp <= transaction.Timestamp {
				continue
			}

			err := database.Client.Put(key, data)

			if err == nil {
				r.Received[key] = transaction.Timestamp
			} else {
				log.Println("Replication error:", err)
			}
		}
	}
}

func (r *Replication) Transmit() {
	var transaction Transaction
	for {
		transaction = <-r.Transmitter
		if len(r.Replicas) == 0 {
			continue
		}
		body, err := proto.Marshal(&transaction)

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
			}
		} else {
			log.Println("Replication error:", err)
		}
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
		Transmitter: make(chan Transaction, 128),
		Receiver:    make(chan Transaction, len(replications)*128),
		Received:    make(map[string]int64),
		Replicas:    hosts,
	}
}

func NewTransaction(data map[string][]byte) Transaction {
	return Transaction{
		Timestamp: time.Now().UnixNano(),
		Data:      data,
	}
}

func init() {
	replicationHosts := strings.Split(os.Getenv("RECON_REPLICATION_HOSTS"), ",")
	Replica = NewReplication(replicationHosts)
	go Replica.Receive()

	if os.Getenv("RECON_REPLICATION_HOSTS") != "" && len(replicationHosts) > 0 {
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
