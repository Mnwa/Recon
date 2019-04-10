package main_test

import (
	"Recon/backup"
	"Recon/database"
	"github.com/prologic/bitcask"
	"log"
	"testing"
)

func InitBackup() {
	var err error
	database.Client, err = bitcask.Open("/tmp/test_recon")
	if err != nil {
		log.Fatal(err)
	}
}

func TestBackup(t *testing.T) {
	InitBackup()
	data, err := backup.CreateBackup()
	if err != nil {
		t.Error(err)
	}
	if len(data) == 0 {
		t.Error("Backup must not be empty")
	}

	err = backup.RestoreBackup(data)
	if err != nil {
		t.Error(err)
	}
	database.Client.Close()
}
