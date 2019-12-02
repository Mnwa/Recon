package controllers

import (
	"Recon/adapters"
	"github.com/valyala/fasthttp"
)

var envAdapter = adapters.NewEnvAdapter()

func CreateEnv(ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()
	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	err := envAdapter.Create(project, projectType, body)
	if err == nil {
		ctx.SetBody(body)
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}
}

func CreateKeyEnv(ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()
	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	key := ctx.UserValue("key").(string)
	err := envAdapter.CreateKey(project, projectType, key, body)
	if err == nil {
		ctx.SetBody(body)
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}
}

func UpdateEnv(ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()
	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	err := envAdapter.Update(project, projectType, body)
	if err == nil {
		data, _ := envAdapter.Get(project, projectType)
		ctx.SetBody(data)
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}
}

func UpdateKeyEnv(ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()
	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	key := ctx.UserValue("key").(string)
	err := envAdapter.UpdateKey(project, projectType, key, body)
	if err == nil {
		ctx.SetBody(body)
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}
}

func GetEnv(ctx *fasthttp.RequestCtx) {
	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	body, err := envAdapter.Get(project, projectType)
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

func GetKeyEnv(ctx *fasthttp.RequestCtx) {
	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	key := ctx.UserValue("key").(string)
	body, err := envAdapter.GetKey(project, projectType, key)
	if err == nil {
		ctx.SetBody(body)
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}
}

func DeleteEnv(ctx *fasthttp.RequestCtx) {
	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	err := envAdapter.Delete(project, projectType)
	if err == nil {
		ctx.SetBodyString("Deleted")
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}
}

func DeleteKeyEnv(ctx *fasthttp.RequestCtx) {
	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	key := ctx.UserValue("key").(string)
	err := envAdapter.DeleteKey(project, projectType, key)
	if err == nil {
		ctx.SetBodyString("Deleted")
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}
}
