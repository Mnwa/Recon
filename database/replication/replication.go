package replication

import (
	"Recon/database"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
	"time"
)

type Message struct {
	Timestamp int64             `json:"timestamp"`
	Data      map[string][]byte `json:"data"`
}

func (m *Message) Keys() []string {
	var keys []string
	for key := range m.Data {
		keys = append(keys, key)
	}
	return keys
}

type Replication struct {
	Transmitter chan Message
	Receiver    chan Message

	Received map[string]int64

	Replicas []*fasthttp.HostClient
}

func (r *Replication) Receive() {
	var message Message
	for {
		message = <-r.Receiver

		for key, data := range message.Data {
			if timestamp, ok := r.Received[key]; ok && timestamp <= message.Timestamp {
				continue
			}

			err := database.Client.Put(key, data)

			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (r *Replication) Transmit() {
	var message Message
	for {
		message = <-r.Transmitter
		body, err := json.Marshal(message)

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
		r.Transmitter <- NewMessage(data)
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
		Transmitter: make(chan Message, 32),
		Receiver:    make(chan Message),
		Received:    make(map[string]int64),
		Replicas:    hosts,
	}
}

func NewMessage(data map[string][]byte) Message {
	return Message{
		Timestamp: time.Now().UnixNano(),
		Data:      data,
	}
}

var Replica *Replication
