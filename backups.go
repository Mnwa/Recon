package main

import (
	"Recon/backup"
	"github.com/valyala/fasthttp"
)

func GetBackup(ctx *fasthttp.RequestCtx) {
	data, err := backup.CreateBackup()
	if err == nil {
		ctx.SetContentType("application/json")
		ctx.SetBody(data)
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}

}
func RestoreBackup(ctx *fasthttp.RequestCtx) {
	err := backup.RestoreBackup(ctx.PostBody())
	if err == nil {
		ctx.SetBodyString("Success restored")
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}
}
