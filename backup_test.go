package main_test

import (
	"Recon/backup"
	"testing"
)

func TestBackup(t *testing.T) {
	data, err := backup.CreateBackup()
	if err != nil {
		t.Error(err)
	}

	err = backup.RestoreBackup(data)
	if err != nil {
		t.Error(err)
	}
}
