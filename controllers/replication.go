package controllers

import (
	"Recon/database/replication"
	"github.com/gogo/protobuf/proto"
	"github.com/valyala/fasthttp"
	"log"
)

func ReceiveMessagesReplication(ctx *fasthttp.RequestCtx) {
	transaction := replication.AcquireTransaction()
	err := proto.Unmarshal(ctx.PostBody(), transaction)
	if err == nil {
		replication.Replica.Receiver <- transaction
	} else {
		log.Println(err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}
}
