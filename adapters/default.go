package adapters

import (
	"Recon/database"
	"Recon/database/replication"
	"bytes"
	"errors"
	"log"
	"strings"
	"sync"
)

type Default struct {
	data map[string][]byte
}

func (e *Default) parseDefault(data []byte) {
	for _, val := range strings.Split(string(data), "\n") {
		if val != "" {
			row := strings.SplitN(val, "=", 2)
			if len(row) == 2 {
				e.data[strings.TrimSpace(row[0])] = []byte(strings.TrimSpace(row[1]))
			}
		}
	}
}

func (e *Default) Create(project string, projectType string, data []byte) error {
	var err error
	replicationData := make(map[string][]byte)
	_ = e.Delete(project, projectType)
	e.parseDefault(data)
	for key, value := range e.data {
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
	var err error
	replicationData := make(map[string][]byte)

	e.parseDefault(data)
	for key, value := range e.data {
		storageKey := strings.ToLower(project + "/" + projectType + "/" + key)
		database.Client.Set(storageKey, value)

		replicationData[storageKey] = value
	}
	go replication.Replica.SendMessage(replicationData)

	return err
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
	if projectType != "default" {
		_, err := e.Get(project, "default")

		if err != nil {
			return nil, err
		}
	}
	storageKey := strings.ToLower(project + "/" + projectType + "/")
	database.Client.Scan(storageKey, func(key string, value []byte) bool {
		// 0x04 byte is removed val
		if !bytes.Equal(value, []byte{0x04}) {
			e.data[strings.ToUpper(strings.ReplaceAll(key, storageKey, ""))] = value
		}
		return true
	})

	for key, value := range e.data {
		data += key + "=" + string(value) + "\n"
	}
	return []byte(data), nil
}

func (e *Default) GetKey(project string, projectType string, key string) ([]byte, error) {
	storageKey := strings.ToLower(project + "/" + projectType + "/" + key)
	value, err := database.Client.Get(storageKey)
	return value, err
}

func (e *Default) Delete(project string, projectType string) error {
	replicationData := make(map[string][]byte)

	storageKey := strings.ToLower(project + "/" + projectType + "/")
	database.Client.Scan(storageKey, func(key string, value []byte) bool {
		err := database.Client.Del(key)
		if err == nil {
			replicationData[storageKey] = nil
			go replication.Replica.SendMessage(replicationData)
			log.Println(err)
			return true
		}

		return false
	})
	return nil
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

func GetDefault() *Default {
	return defaultPool.Get().(*Default)
}

func PutDefault(e *Default) {
	e.data = make(map[string][]byte)

	defaultPool.Put(e)
}

var defaultPool = sync.Pool{
	New: func() interface{} {
		return &Default{
			data: make(map[string][]byte),
		}
	},
}
