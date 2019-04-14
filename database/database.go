package database

import (
	"github.com/prologic/bitcask"
	"log"
	"os"
)

var Client *bitcask.Bitcask

func init() {
	var err error
	dbDir := os.Getenv("RECON_DB_DIR")
	if dbDir == "" {
		dbDir = "/tmp/test_recon"
	}

	Client, err = bitcask.Open(dbDir)
	if err != nil {
		log.Fatal(err)
	}
}
