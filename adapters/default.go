package adapters

import (
	"Recon/database"
	"Recon/database/replication"
	"bytes"
	"errors"
	"strings"
)

type Default struct{}

func (e *Default) Create(project string, projectType string, data []byte) error {
	replicationData := make(map[string][]byte)
	err := e.Delete(project, projectType)
	if err != nil {
		return err
	}
	parsedData := parseDefault(data)
	for key, value := range parsedData {
		storageKey := strings.ToLower(project + "/" + projectType + "/" + key)
		database.Client.Set(storageKey, value)
		replicationData[storageKey] = value
	}
	go replication.Replica.SendMessage(replicationData)

	return err
}

func (e *Default) CreateKey(project string, projectType string, key string, data []byte) error {
	replicationData := make(map[string][]byte)
	storageKey := strings.ToLower(project + "/" + projectType + "/" + key)

	database.Client.Set(storageKey, data)

	replicationData[storageKey] = data
	go replication.Replica.SendMessage(replicationData)

	return nil
}

func (e *Default) Update(project string, projectType string, data []byte) error {
	replicationData := make(map[string][]byte)

	parsedData := parseDefault(data)
	for key, value := range parsedData {
		storageKey := strings.ToLower(project + "/" + projectType + "/" + key)
		database.Client.Set(storageKey, value)

		replicationData[storageKey] = value
	}
	go replication.Replica.SendMessage(replicationData)

	return nil
}

func (e *Default) UpdateKey(project string, projectType string, key string, data []byte) error {
	replicationData := make(map[string][]byte)
	storageKey := strings.ToLower(project + "/" + projectType + "/" + key)
	_, err := database.Client.Get(storageKey)

	if err == nil {
		database.Client.Set(storageKey, data)

		replicationData[storageKey] = data
		go replication.Replica.SendMessage(replicationData)

		return err
	} else {
		return errors.New("error: key not found")
	}
}

func (e *Default) Get(project string, projectType string) ([]byte, error) {
	var data = ""
	result := make(map[string][]byte)
	if projectType != "default" {
		defaultData, err := e.Get(project, "default")

		if err != nil {
			return nil, err
		}

		result = parseDefault(defaultData)
	}
	storageKey := strings.ToLower(project + "/" + projectType + "/")
	database.Client.Scan(storageKey, func(key string, value []byte) bool {
		// 0x04 byte is removed value marker
		if !bytes.Equal(value, []byte{0x04}) {
			result[strings.ReplaceAll(key, storageKey, "")] = value
		}
		return true
	})

	for key, value := range result {
		data += key + "=" + string(value) + "\n"
	}
	return []byte(data), nil
}

func (e *Default) GetKey(project string, projectType string, key string) ([]byte, error) {
	storageKey := strings.ToLower(project + "/" + projectType + "/" + key)
	value, err := database.Client.Get(storageKey)
	return value, err
}

func (e *Default) Delete(project string, projectType string) (err error) {
	replicationData := make(map[string][]byte)

	storageKey := strings.ToLower(project + "/" + projectType + "/")
	database.Client.Scan(storageKey, func(key string, value []byte) bool {
		err = database.Client.Del(key)
		if err == nil {
			replicationData[storageKey] = nil
			go replication.Replica.SendMessage(replicationData)
			return true
		}

		return false
	})
	return
}

func (e *Default) DeleteKey(project string, projectType string, key string) error {
	replicationData := make(map[string][]byte)
	storageKey := strings.ToLower(project + "/" + projectType + "/" + key)
	err := database.Client.Del(storageKey)

	if err == nil {
		replicationData[storageKey] = nil
		go replication.Replica.SendMessage(replicationData)
	}

	return err
}

func NewDefaultAdapter() Adapter {
	return &Default{}
}

func parseDefault(data []byte) map[string][]byte {
	result := make(map[string][]byte)
	for _, val := range strings.Split(string(data), "\n") {
		if val != "" {
			row := strings.SplitN(val, "=", 2)
			if len(row) == 2 {
				result[strings.TrimSpace(row[0])] = []byte(strings.TrimSpace(row[1]))
			}
		}
	}
	return result
}
