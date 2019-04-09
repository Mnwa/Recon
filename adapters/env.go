package adapters

import (
	"Recon/database"
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
			e.data[row[0]] = []byte(row[1])
		}
	}
}

func (e *Env) Create(project string, projectType string, data []byte) error {
	var err error
	_ = e.Delete(project, projectType)
	e.parseEnv(data)
	for key, value := range e.data {
		storageKey := strings.ToLower(project + "/" + projectType + "/" + key)
		err = database.Client.Put(storageKey, value)

		if err != nil {
			break
		}
	}
	return err
}

func (e *Env) CreateKey(project string, projectType string, key string, data []byte) error {
	storageKey := strings.ToLower(project + "/" + projectType + "/" + key)

	return database.Client.Put(storageKey, data)
}

func (e *Env) Update(project string, projectType string, data []byte) error {
	var oldData, err = e.Get(project, projectType)
	if err != nil {
		return nil
	}

	e.parseEnv(oldData)
	e.parseEnv(data)
	for key, value := range e.data {
		storageKey := strings.ToLower(project + "/" + projectType + "/" + key)
		err = database.Client.Put(storageKey, value)

		if err != nil {
			break
		}
	}
	return err
}

func (e *Env) UpdateKey(project string, projectType string, key string, data []byte) error {
	storageKey := strings.ToLower(project + "/" + projectType + "/" + key)

	if database.Client.Has(storageKey) {
		return database.Client.Put(storageKey, data)
	} else {
		return errors.New("error: key not found")
	}
}

func (e *Env) Get(project string, projectType string) ([]byte, error) {
	var data = ""
	storageKey := strings.ToLower(project + "/" + projectType + "/")
	err := database.Client.Scan(storageKey, func(key string) error {
		if database.Client.Has(key) {
			value, err := database.Client.Get(key)
			data += strings.ToUpper(strings.ReplaceAll(key, storageKey, "")) + "=" + strings.ToUpper(string(value)) + "\n"
			return err
		} else {
			return nil
		}
	})
	return []byte(data), err
}

func (e *Env) GetKey(project string, projectType string, key string) ([]byte, error) {
	storageKey := strings.ToLower(project + "/" + projectType + "/" + key)
	value, err := database.Client.Get(storageKey)
	return bytes.ToUpper(value), err
}

func (e *Env) Delete(project string, projectType string) error {
	storageKey := strings.ToLower(project + "/" + projectType + "/")
	err := database.Client.Scan(storageKey, func(key string) error {
		if database.Client.Has(key) {
			err := database.Client.Delete(key)
			return err
		} else {
			return nil
		}
	})
	return err
}

func (e *Env) DeleteKey(project string, projectType string, key string) error {
	storageKey := strings.ToLower(project + "/" + projectType + "/" + key)
	return database.Client.Delete(storageKey)
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
