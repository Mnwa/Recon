package main

import (
	"Recon/adapters"
	"github.com/valyala/fasthttp"
)

func CreateDefault(ctx *fasthttp.RequestCtx) {
	var adapter = adapters.GetDefault()

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

	adapters.PutDefault(adapter)
}

func CreateKeyDefault(ctx *fasthttp.RequestCtx) {
	var adapter = adapters.GetDefault()

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

	adapters.PutDefault(adapter)
}

func UpdateDefault(ctx *fasthttp.RequestCtx) {
	var adapter = adapters.GetDefault()

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

	adapters.PutDefault(adapter)
}

func UpdateKeyDefault(ctx *fasthttp.RequestCtx) {
	var adapter = adapters.GetDefault()

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

	adapters.PutDefault(adapter)
}

func GetDefault(ctx *fasthttp.RequestCtx) {
	var adapter = adapters.GetDefault()

	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	body, err := adapter.Get(project, projectType)
	if err == nil {
		ctx.SetBody(body)
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}

	adapters.PutDefault(adapter)
}

func GetKeyDefault(ctx *fasthttp.RequestCtx) {
	var adapter = adapters.GetDefault()

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

	adapters.PutDefault(adapter)
}

func DeleteDefault(ctx *fasthttp.RequestCtx) {
	var adapter = adapters.GetDefault()

	project := ctx.UserValue("project").(string)
	projectType := ctx.UserValue("type").(string)
	err := adapter.Delete(project, projectType)
	if err == nil {
		ctx.SetBodyString("Deleted")
	} else {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
	}

	adapters.PutDefault(adapter)
}

func DeleteKeyDefault(ctx *fasthttp.RequestCtx) {
	var adapter = adapters.GetDefault()

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

	adapters.PutDefault(adapter)
}
