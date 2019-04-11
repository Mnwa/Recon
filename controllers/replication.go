package controllers

import (
	"Recon/database/replication"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
)

func RecieveMessagesReplication(ctx *fasthttp.RequestCtx) {
	var message replication.Message
	err := json.Unmarshal(ctx.PostBody(), &message)
	if err == nil {
		replication.Replica.Receiver <- message
	} else {
		log.Println(err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}
}
