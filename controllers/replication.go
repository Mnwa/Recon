package controllers

import (
	"Recon/database/replication"
	"github.com/gogo/protobuf/proto"
	"github.com/valyala/fasthttp"
	"log"
)

func RecieveMessagesReplication(ctx *fasthttp.RequestCtx) {
	transaction := replication.GetTransaction()
	err := proto.Unmarshal(ctx.PostBody(), transaction)
	if err == nil {
		replication.Replica.Receiver <- transaction
	} else {
		log.Println(err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}
}
