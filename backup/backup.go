package backup

import (
	"Recon/database"
	"github.com/golang/protobuf/proto"
)

func CreateBackup() ([]byte, error) {
	var backup = &Backup{
		Data: make(map[string][]byte),
	}
	database.Client.Scan("", func(key string, value []byte) bool {
		backup.Data[key] = value
		return true
	})
	return proto.Marshal(backup)
}

func RestoreBackup(body []byte) error {
	var backup Backup
	err := proto.Unmarshal(body, &backup)
	if err == nil {
		for key, value := range backup.GetData() {
			database.Client.Set(key, value)
		}
		return nil
	} else {
		return err
	}
}
