package controllers

import (
	"Recon/adapters"
	"github.com/valyala/fasthttp"
)

func CreateEnv(ctx *fasthttp.RequestCtx) {
	var adapter = adapters.GetEnv()

	body := ctx.PostBody()
	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	err := adapter.Create(project, projectType, body)
	if err == nil {
		ctx.SetBody(body)
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}

	adapters.PutEnv(adapter)
}

func CreateKeyEnv(ctx *fasthttp.RequestCtx) {
	var adapter = adapters.GetEnv()

	body := ctx.PostBody()
	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	key := ctx.UserValue("key").(string)
	err := adapter.CreateKey(project, projectType, key, body)
	if err == nil {
		ctx.SetBody(body)
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}

	adapters.PutEnv(adapter)
}

func UpdateEnv(ctx *fasthttp.RequestCtx) {
	var adapter = adapters.GetEnv()

	body := ctx.PostBody()
	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	err := adapter.Update(project, projectType, body)
	if err == nil {
		data, _ := adapter.Get(project, projectType)
		ctx.SetBody(data)
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}

	adapters.PutEnv(adapter)
}

func UpdateKeyEnv(ctx *fasthttp.RequestCtx) {
	var adapter = adapters.GetEnv()

	body := ctx.PostBody()
	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	key := ctx.UserValue("key").(string)
	err := adapter.UpdateKey(project, projectType, key, body)
	if err == nil {
		ctx.SetBody(body)
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}

	adapters.PutEnv(adapter)
}

func GetEnv(ctx *fasthttp.RequestCtx) {
	var adapter = adapters.GetEnv()

	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	body, err := adapter.Get(project, projectType)
	if err == nil {
		ctx.SetBody(body)
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}

	adapters.PutEnv(adapter)
}

func GetKeyEnv(ctx *fasthttp.RequestCtx) {
	var adapter = adapters.GetEnv()

	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	key := ctx.UserValue("key").(string)
	body, err := adapter.GetKey(project, projectType, key)
	if err == nil {
		ctx.SetBody(body)
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}

	adapters.PutEnv(adapter)
}

func DeleteEnv(ctx *fasthttp.RequestCtx) {
	var adapter = adapters.GetEnv()

	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	err := adapter.Delete(project, projectType)
	if err == nil {
		ctx.SetBodyString("Deleted")
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}

	adapters.PutEnv(adapter)
}

func DeleteKeyEnv(ctx *fasthttp.RequestCtx) {
	var adapter = adapters.GetEnv()

	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	key := ctx.UserValue("key").(string)
	err := adapter.DeleteKey(project, projectType, key)
	if err == nil {
		ctx.SetBodyString("Deleted")
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}

	adapters.PutEnv(adapter)
}
