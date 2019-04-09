package main

import (
	"Recon/database"
	"fmt"
	"github.com/prologic/bitcask"
	"log"
	"os"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func main() {
	var err error
	dbDir := os.Getenv("RECON_DB_DIR")
	if dbDir == "" {
		dbDir = "/var/lib/recon"
	}

	addr := os.Getenv("RECON_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	database.Client, err = bitcask.Open(dbDir)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Client.Close()
	router := fasthttprouter.New()
	router.GET("/", Index)

	// Env based
	router.GET("/:project/:type/env", GetEnv)
	router.PUT("/:project/:type/env", CreateEnv)
	router.POST("/:project/:type/env", UpdateEnv)
	router.DELETE("/:project/:type/env", DeleteEnv)

	router.GET("/:project/:type/env/:key", GetKeyEnv)
	router.PUT("/:project/:type/env/:key", CreateKeyEnv)
	router.POST("/:project/:type/env/:key", UpdateKeyEnv)
	router.DELETE("/:project/:type/env/:key", DeleteKeyEnv)

	// Default kv
	router.GET("/:project/:type/config", GetDefault)
	router.PUT("/:project/:type/config", CreateDefault)
	router.POST("/:project/:type/config", UpdateDefault)
	router.DELETE("/:project/:type/config", DeleteDefault)

	router.GET("/:project/:type/config/:key", GetKeyDefault)
	router.PUT("/:project/:type/config/:key", CreateKeyDefault)
	router.POST("/:project/:type/config/:key", UpdateKeyDefault)
	router.DELETE("/:project/:type/config/:key", DeleteKeyDefault)

	log.Println("Recon started")

	log.Fatal(fasthttp.ListenAndServe(addr, router.Handler))
}
