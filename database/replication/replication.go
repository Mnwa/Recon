package replication

import (
	"Recon/database"
	"github.com/gogo/protobuf/proto"
	"github.com/valyala/fasthttp"
	"log"
	"time"
)

type Replication struct {
	Transmitter chan Transaction
	Receiver    chan Transaction

	Received map[string]int64

	Replicas []*fasthttp.HostClient
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

			if err != nil {
				log.Println(err)
			} else {
				r.Received[key] = transaction.Timestamp
			}
		}
	}
}

func (r *Replication) Transmit() {
	var transaction Transaction
	for {
		transaction = <-r.Transmitter
		body, err := proto.Marshal(&transaction)

		if err == nil {
			for _, host := range r.Replicas {
				req := fasthttp.AcquireRequest()
				req.SetRequestURI("/replication/receiver")
				req.Header.SetMethod("POST")
				req.SetBody(body)
				resp := fasthttp.AcquireResponse()

				err := host.Do(req, resp)

				if err != nil {
					log.Println(err)
				}
			}
		} else {
			log.Println(err)
		}
	}
}

func (r *Replication) SendMessage(data map[string][]byte) {
	if len(r.Replicas) > 0 {
		r.Transmitter <- NewTransaction(data)
	}
}

func NewReplication(replications []string) *Replication {
	var hosts []*fasthttp.HostClient
	for _, addr := range replications {
		hosts = append(hosts, &fasthttp.HostClient{
			Addr: addr,
		})
	}
	return &Replication{
		Transmitter: make(chan Transaction, 32),
		Receiver:    make(chan Transaction),
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

var Replica *Replication
