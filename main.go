package main

import (
	"Recon/database"
	"github.com/prologic/bitcask"
	"log"
	"os"

	"github.com/buaazp/fasthttprouter"
	fastp "github.com/flf2ko/fasthttp-prometheus"
	"github.com/valyala/fasthttp"
)

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

	router.GET("/backup", GetBackup)
	router.POST("/backup", RestoreBackup)

	// Env based
	router.GET("/projects/:project/:type/env", GetEnv)
	router.PUT("/projects/:project/:type/env", CreateEnv)
	router.POST("/projects/:project/:type/env", UpdateEnv)
	router.DELETE("/projects/:project/:type/env", DeleteEnv)

	router.GET("/projects/:project/:type/env/:key", GetKeyEnv)
	router.PUT("/projects/:project/:type/env/:key", CreateKeyEnv)
	router.POST("/projects/:project/:type/env/:key", UpdateKeyEnv)
	router.DELETE("/projects/:project/:type/env/:key", DeleteKeyEnv)

	// Default kv
	router.GET("/projects/:project/:type/config", GetDefault)
	router.PUT("/projects/:project/:type/config", CreateDefault)
	router.POST("/projects/:project/:type/config", UpdateDefault)
	router.DELETE("/projects/:project/:type/config", DeleteDefault)

	router.GET("/projects/:project/:type/config/:key", GetKeyDefault)
	router.PUT("/projects/:project/:type/config/:key", CreateKeyDefault)
	router.POST("/projects/:project/:type/config/:key", UpdateKeyDefault)
	router.DELETE("/projects/:project/:type/config/:key", DeleteKeyDefault)

	p := fastp.NewPrometheus("recon")
	p.MetricsPath = "/"
	fastpHandler := p.WrapHandler(router)

	log.Println("Recon started")

	log.Fatal(fasthttp.ListenAndServe(addr, fastpHandler))
}
