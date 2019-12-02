package controllers

import (
	"Recon/adapters"
	"github.com/valyala/fasthttp"
)

var defaultAdapter = adapters.NewDefaultAdapter()

func CreateDefault(ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()
	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	err := defaultAdapter.Create(project, projectType, body)
	if err == nil {
		ctx.SetBody(body)
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}
}

func CreateKeyDefault(ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()
	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	key := ctx.UserValue("key").(string)
	err := defaultAdapter.CreateKey(project, projectType, key, body)
	if err == nil {
		ctx.SetBody(body)
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}
}

func UpdateDefault(ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()
	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	err := defaultAdapter.Update(project, projectType, body)
	if err == nil {
		data, _ := defaultAdapter.Get(project, projectType)
		ctx.SetBody(data)
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}
}

func UpdateKeyDefault(ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()
	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	key := ctx.UserValue("key").(string)
	err := defaultAdapter.UpdateKey(project, projectType, key, body)
	if err == nil {
		ctx.SetBody(body)
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}
}

func GetDefault(ctx *fasthttp.RequestCtx) {
	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	body, err := defaultAdapter.Get(project, projectType)
	if err == nil {
		if ctx.Request.Header.HasAcceptEncoding("gzip") {
			ctx.Response.Header.Add("Content-Encoding", "gzip")
			ctx.SetBody(fasthttp.AppendGzipBytes(nil, body))
		} else {
			ctx.SetBody(body)
		}
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}
}

func GetKeyDefault(ctx *fasthttp.RequestCtx) {
	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	key := ctx.UserValue("key").(string)
	body, err := defaultAdapter.GetKey(project, projectType, key)
	if err == nil {
		ctx.SetBody(body)
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}
}

func DeleteDefault(ctx *fasthttp.RequestCtx) {
	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	err := defaultAdapter.Delete(project, projectType)
	if err == nil {
		ctx.SetBodyString("Deleted")
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}
}

func DeleteKeyDefault(ctx *fasthttp.RequestCtx) {
	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	key := ctx.UserValue("key").(string)
	err := defaultAdapter.DeleteKey(project, projectType, key)
	if err == nil {
		ctx.SetBodyString("Deleted")
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}
}
