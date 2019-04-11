package backup

import (
	"Recon/database"
	"github.com/gogo/protobuf/proto"
)

func CreateBackup() ([]byte, error) {
	var backup = &Backup{
		Data: make(map[string][]byte),
	}
	err := database.Client.Fold(func(key string) error {
		if database.Client.Has(key) {
			value, err := database.Client.Get(key)
			backup.Data[key] = value
			return err
		} else {
			return nil
		}
	})
	if err == nil {
		return proto.Marshal(backup)
	} else {
		return nil, err
	}
}

func RestoreBackup(body []byte) error {
	var backup Backup
	err := proto.Unmarshal(body, &backup)
	if err == nil {
		for key, value := range backup.GetData() {
			err = database.Client.Put(key, value)
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		return err
	}
}
