package database

import (
	"github.com/Mnwa/ReconEngine"
	"os"
)

var Client reconEngine.MemStorage

func init() {
	dbDir := os.Getenv("RECON_DB_DIR")
	if dbDir == "" {
		dbDir = os.TempDir() + "/test_recon"
	}
	_ = os.Mkdir(dbDir, 0750)

	Client = reconEngine.NewMem(nil, &dbDir)
}
