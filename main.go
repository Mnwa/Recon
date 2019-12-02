package main

import (
	"Recon/controllers"
	"Recon/database"
	fastp "github.com/Mnwa/fasthttprouter-prometheus"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"time"
)

func main() {
	addr := os.Getenv("RECON_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	go func() {
		// save data to disk every second
		for {
			if database.Client.Len() > 0 {
				log.Println("Save data to partitions")
				err := database.Client.Sync()
				if err != nil {
					log.Fatal(err)
				}
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		// run merge sort every hour
		for {
			if database.Client.SsTable().Len() > 2 {
				log.Println("Merging partitions..")
				err := database.Client.SsTable().MergeSort()
				if err != nil {
					log.Fatal(err)
				}
			}
			time.Sleep(time.Hour)
		}
	}()

	defer database.Client.Sync()
	r := router.New()

	r.GET("/backup", controllers.GetBackup)
	r.POST("/backup", controllers.RestoreBackup)

	r.POST("/replication/receiver", controllers.ReceiveMessagesReplication)

	// Env based
	r.GET("/projects/:project/:type/env", controllers.GetEnv)
	r.PUT("/projects/:project/:type/env", controllers.CreateEnv)
	r.POST("/projects/:project/:type/env", controllers.UpdateEnv)
	r.DELETE("/projects/:project/:type/env", controllers.DeleteEnv)

	r.GET("/projects/:project/:type/env/:key", controllers.GetKeyEnv)
	r.PUT("/projects/:project/:type/env/:key", controllers.CreateKeyEnv)
	r.POST("/projects/:project/:type/env/:key", controllers.UpdateKeyEnv)
	r.DELETE("/projects/:project/:type/env/:key", controllers.DeleteKeyEnv)

	// Default kv
	r.GET("/projects/:project/:type/config", controllers.GetDefault)
	r.PUT("/projects/:project/:type/config", controllers.CreateDefault)
	r.POST("/projects/:project/:type/config", controllers.UpdateDefault)
	r.DELETE("/projects/:project/:type/config", controllers.DeleteDefault)

	r.GET("/projects/:project/:type/config/:key", controllers.GetKeyDefault)
	r.PUT("/projects/:project/:type/config/:key", controllers.CreateKeyDefault)
	r.POST("/projects/:project/:type/config/:key", controllers.UpdateKeyDefault)
	r.DELETE("/projects/:project/:type/config/:key", controllers.DeleteKeyDefault)

	p := fastp.NewPrometheus("recon")
	p.MetricsPath = "/"
	fastpHandler := p.WrapHandler(r)

	log.Println("Recon started")

	log.Fatal(fasthttp.ListenAndServe(addr, fastpHandler))
}
