package adapters

import (
	"Recon/database"
	"Recon/database/replication"
	"bytes"
	"errors"
	"strings"
	"sync"
)

type Env struct {
	data map[string][]byte
}

func (e *Env) parseEnv(data []byte) {
	for _, val := range strings.Split(string(data), "\n") {
		if val != "" {
			row := strings.Split(val, "=")
			if len(row) == 2 {
				e.data[row[0]] = []byte(row[1])
			}
		}
	}
}

func (e *Env) Create(project string, projectType string, data []byte) error {
	var err error
	replicationData := make(map[string][]byte)
	_ = e.Delete(project, projectType)
	e.parseEnv(data)
	for key, value := range e.data {
		storageKey := strings.ToLower(project + "/" + projectType + "/" + key)
		err = database.Client.Put(storageKey, value)

		if err != nil {
			break
		}
		replicationData[storageKey] = value
	}
	go replication.Replica.SendMessage(replicationData)
	return err
}

func (e *Env) CreateKey(project string, projectType string, key string, data []byte) error {
	replicationData := make(map[string][]byte)
	storageKey := strings.ToLower(project + "/" + projectType + "/" + key)

	err := database.Client.Put(storageKey, data)

	if err == nil {
		replicationData[storageKey] = data
		go replication.Replica.SendMessage(replicationData)
	}

	return err
}

func (e *Env) Update(project string, projectType string, data []byte) error {
	var err error
	replicationData := make(map[string][]byte)

	e.parseEnv(data)
	for key, value := range e.data {
		storageKey := strings.ToLower(project + "/" + projectType + "/" + key)
		err := database.Client.Put(storageKey, value)

		if err != nil {
			break
		}
		replicationData[storageKey] = value
	}
	go replication.Replica.SendMessage(replicationData)
	return err
}

func (e *Env) UpdateKey(project string, projectType string, key string, data []byte) error {
	replicationData := make(map[string][]byte)
	storageKey := strings.ToLower(project + "/" + projectType + "/" + key)

	if database.Client.Has(storageKey) {
		err := database.Client.Put(storageKey, data)

		if err == nil {
			replicationData[storageKey] = data
			go replication.Replica.SendMessage(replicationData)
		}

		return err
	} else {
		return errors.New("error: key not found")
	}
}

func (e *Env) Get(project string, projectType string) ([]byte, error) {
	var data = ""
	if projectType != "default" {
		_, err := e.Get(project, "default")

		if err != nil {
			return nil, err
		}
	}
	storageKey := strings.ToLower(project + "/" + projectType + "/")
	err := database.Client.Scan(storageKey, func(key string) error {
		if database.Client.Has(key) {
			value, err := database.Client.Get(key)
			if err == nil {
				e.data[strings.ToUpper(strings.ReplaceAll(key, storageKey, ""))] = value
			}
			return err
		} else {
			return nil
		}
	})

	for key, value := range e.data {
		data += key + "=" + string(value) + "\n"
	}
	return []byte(data), err
}

func (e *Env) GetKey(project string, projectType string, key string) ([]byte, error) {
	storageKey := strings.ToLower(project + "/" + projectType + "/" + key)
	value, err := database.Client.Get(storageKey)
	return bytes.ToUpper(value), err
}

func (e *Env) Delete(project string, projectType string) error {
	replicationData := make(map[string][]byte)

	storageKey := strings.ToLower(project + "/" + projectType + "/")
	err := database.Client.Scan(storageKey, func(key string) error {
		if database.Client.Has(key) {
			err := database.Client.Delete(key)
			if err == nil {
				replicationData[storageKey] = nil
				go replication.Replica.SendMessage(replicationData)
			}
			return err
		} else {
			return nil
		}
	})
	return err
}

func (e *Env) DeleteKey(project string, projectType string, key string) error {
	replicationData := make(map[string][]byte)
	storageKey := strings.ToLower(project + "/" + projectType + "/" + key)
	err := database.Client.Delete(storageKey)

	if err == nil {
		replicationData[storageKey] = nil
		go replication.Replica.SendMessage(replicationData)
	}

	return err
}

func GetEnv() *Env {
	return envPool.Get().(*Env)
}

func PutEnv(e *Env) {
	e.data = make(map[string][]byte)

	envPool.Put(e)
}

var envPool = sync.Pool{
	New: func() interface{} {
		return &Env{
			data: make(map[string][]byte),
		}
	},
}
