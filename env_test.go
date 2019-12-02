package main_test

import (
	"Recon/adapters"
	"Recon/database"
	"testing"
)

func TestEnvCreate(t *testing.T) {
	var adapter = adapters.NewEnvAdapter()
	err := adapter.Create("testProject", "testType", []byte("VAR_TEST=1"))
	if err != nil {
		t.Error(err)
	}
	if database.Client.Len() == 0 {
		t.Error("Not created")
	}
}

func TestEnvCreateKey(t *testing.T) {
	var adapter = adapters.NewEnvAdapter()
	err := adapter.CreateKey("testProject", "testType", "VAR_TEST_TWO", []byte("2"))
	if err != nil {
		t.Error(err)
	}
	if database.Client.Len() == 0 {
		t.Error("Not created key")
	}
}

func TestEnvUpdate(t *testing.T) {
	var adapter = adapters.NewEnvAdapter()
	err := adapter.Update("testProject", "testType", []byte("VAR_TEST=2"))
	if err != nil {
		t.Error(err)
	}
	if database.Client.Len() == 0 {
		t.Error("Not updated")
	}
}

func TestEnvUpdateKey(t *testing.T) {
	var adapter = adapters.NewEnvAdapter()
	err := adapter.UpdateKey("testProject", "testType", "VAR_TEST_TWO", []byte("3"))
	if err != nil {
		t.Error(err)
	}
	if database.Client.Len() == 0 {
		t.Error("Not updated key")
	}
}

func TestEnvGet(t *testing.T) {
	var adapter = adapters.NewEnvAdapter()
	data, err := adapter.Get("testProject", "testType")
	if err != nil {
		t.Error(err)
	}
	if string(data) == "" {
		t.Error("Data must be VAR_TEST=2\nVAR_TEST_TWO=3")
	}
}

func TestEnvGetKey(t *testing.T) {
	var adapter = adapters.NewEnvAdapter()
	data, err := adapter.GetKey("testProject", "testType", "VAR_TEST_TWO")
	if err != nil {
		t.Error(err)
	}
	if string(data) != "3" {
		t.Error("Data must be 3.", string(data), "returned")
	}
}

func TestEnvDelete(t *testing.T) {
	var adapter = adapters.NewEnvAdapter()
	err := adapter.Delete("testProject", "testType")
	if err != nil {
		t.Error(err)
	}
}

func TestEnvDeleteKey(t *testing.T) {
	var adapter = adapters.NewEnvAdapter()
	err := adapter.CreateKey("testProject", "testType", "VAR_TEST_TWO", []byte("2"))
	if err != nil {
		t.Error(err)
	}

	err = adapter.DeleteKey("testProject", "testType", "VAR_TEST_TWO")
	if err != nil {
		t.Error(err)
	}
}
