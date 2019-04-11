package main

import (
	"Recon/controllers"
	"Recon/database"
	"Recon/database/replication"
	"github.com/buaazp/fasthttprouter"
	fastp "github.com/flf2ko/fasthttp-prometheus"
	"github.com/prologic/bitcask"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"strings"
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

	replicationHosts := strings.Split(os.Getenv("RECON_REPLICATION_HOSTS"), ",")
	replication.Replica = replication.NewReplication(replicationHosts)
	go replication.Replica.Receive()

	if len(replicationHosts) > 0 {
		go replication.Replica.Transmit()
	}

	router.GET("/backup", controllers.GetBackup)
	router.POST("/backup", controllers.RestoreBackup)

	router.POST("/replication/receiver", controllers.RecieveMessagesReplication)

	// Env based
	router.GET("/projects/:project/:type/env", controllers.GetEnv)
	router.PUT("/projects/:project/:type/env", controllers.CreateEnv)
	router.POST("/projects/:project/:type/env", controllers.UpdateEnv)
	router.DELETE("/projects/:project/:type/env", controllers.DeleteEnv)

	router.GET("/projects/:project/:type/env/:key", controllers.GetKeyEnv)
	router.PUT("/projects/:project/:type/env/:key", controllers.CreateKeyEnv)
	router.POST("/projects/:project/:type/env/:key", controllers.UpdateKeyEnv)
	router.DELETE("/projects/:project/:type/env/:key", controllers.DeleteKeyEnv)

	// Default kv
	router.GET("/projects/:project/:type/config", controllers.GetDefault)
	router.PUT("/projects/:project/:type/config", controllers.CreateDefault)
	router.POST("/projects/:project/:type/config", controllers.UpdateDefault)
	router.DELETE("/projects/:project/:type/config", controllers.DeleteDefault)

	router.GET("/projects/:project/:type/config/:key", controllers.GetKeyDefault)
	router.PUT("/projects/:project/:type/config/:key", controllers.CreateKeyDefault)
	router.POST("/projects/:project/:type/config/:key", controllers.UpdateKeyDefault)
	router.DELETE("/projects/:project/:type/config/:key", controllers.DeleteKeyDefault)

	p := fastp.NewPrometheus("recon")
	p.MetricsPath = "/"
	fastpHandler := p.WrapHandler(router)

	log.Println("Recon started")

	log.Fatal(fasthttp.ListenAndServe(addr, fastpHandler))
}
