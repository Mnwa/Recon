package backup

import (
	"Recon/database"
	"encoding/json"
)

func CreateBackup() ([]byte, error) {
	var data = make(map[string]string)
	err := database.Client.Fold(func(key string) error {
		if database.Client.Has(key) {
			value, err := database.Client.Get(key)
			data[key] = string(value)
			return err
		} else {
			return nil
		}
	})
	if err == nil {
		return json.Marshal(data)
	} else {
		return nil, err
	}
}

func RestoreBackup(body []byte) error {
	var data map[string]string
	err := json.Unmarshal(body, &data)
	if err == nil {
		for key, value := range data {
			err = database.Client.Put(key, []byte(value))
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		return err
	}
}
