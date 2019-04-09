package main

import (
	"Recon/database"
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func main() {
	defer database.Client.Close()
	router := fasthttprouter.New()
	router.GET("/", Index)

	router.GET("/:project/:type/env", GetEnv)
	router.PUT("/:project/:type/env", CreateEnv)
	router.POST("/:project/:type/env", UpdateEnv)
	router.DELETE("/:project/:type/env", DeleteEnv)

	router.GET("/:project/:type/env/:key", GetKeyEnv)
	router.PUT("/:project/:type/env/:key", CreateKeyEnv)
	router.POST("/:project/:type/env/:key", UpdateKeyEnv)
	router.DELETE("/:project/:type/env/:key", DeleteKeyEnv)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
