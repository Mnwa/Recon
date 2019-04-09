package adapters

import (
	"Recon/database"
	"errors"
	"strings"
	"sync"
)

type Default struct {
	data map[string][]byte
}

func (e *Default) parseDefault(data []byte) {
	for _, val := range strings.Split(string(data), "\n") {
		if val != "" {
			row := strings.Split(val, "=")
			e.data[row[0]] = []byte(row[1])
		}
	}
}

func (e *Default) Create(project string, projectType string, data []byte) error {
	var err error
	_ = e.Delete(project, projectType)
	e.parseDefault(data)
	for key, value := range e.data {
		storageKey := strings.ToLower(project + "/" + projectType + "/" + key)
		err = database.Client.Put(storageKey, value)

		if err != nil {
			break
		}
	}
	return err
}

func (e *Default) CreateKey(project string, projectType string, key string, data []byte) error {
	storageKey := strings.ToLower(project + "/" + projectType + "/" + key)

	return database.Client.Put(storageKey, data)
}

func (e *Default) Update(project string, projectType string, data []byte) error {
	var err error

	e.parseDefault(data)
	for key, value := range e.data {
		storageKey := strings.ToLower(project + "/" + projectType + "/" + key)
		err = database.Client.Put(storageKey, value)

		if err != nil {
			break
		}
	}
	return err
}

func (e *Default) UpdateKey(project string, projectType string, key string, data []byte) error {
	storageKey := strings.ToLower(project + "/" + projectType + "/" + key)

	if database.Client.Has(storageKey) {
		return database.Client.Put(storageKey, data)
	} else {
		return errors.New("error: key not found")
	}
}

func (e *Default) Get(project string, projectType string) ([]byte, error) {
	var data = ""
	if projectType != "default" {
		defaultData, err := e.Get(project, "default")
		data = string(defaultData)

		if err != nil {
			return nil, err
		}
	}
	storageKey := strings.ToLower(project + "/" + projectType + "/")
	err := database.Client.Scan(storageKey, func(key string) error {
		if database.Client.Has(key) {
			value, err := database.Client.Get(key)
			data += strings.ReplaceAll(key, storageKey, "") + "=" + string(value) + "\n"
			return err
		} else {
			return nil
		}
	})
	return []byte(data), err
}

func (e *Default) GetKey(project string, projectType string, key string) ([]byte, error) {
	storageKey := strings.ToLower(project + "/" + projectType + "/" + key)
	value, err := database.Client.Get(storageKey)
	return value, err
}

func (e *Default) Delete(project string, projectType string) error {
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

func (e *Default) DeleteKey(project string, projectType string, key string) error {
	storageKey := strings.ToLower(project + "/" + projectType + "/" + key)
	return database.Client.Delete(storageKey)
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
