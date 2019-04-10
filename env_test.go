package main_test

import (
	"Recon/adapters"
	"Recon/database"
	"github.com/prologic/bitcask"
	"log"
	"testing"
)

func InitEnv() *adapters.Env {
	var err error
	database.Client, err = bitcask.Open("/tmp/test_recon")
	if err != nil {
		log.Fatal(err)
	}

	return adapters.GetEnv()
}

func TestEnvCreate(t *testing.T) {
	var adapter = InitEnv()
	err := adapter.Create("testProject", "testType", []byte("VAR_TEST=1"))
	if err != nil {
		t.Error(err)
	}
	database.Client.Close()
}

func TestEnvCreateKey(t *testing.T) {
	var adapter = InitEnv()
	err := adapter.CreateKey("testProject", "testType", "VAR_TEST_TWO", []byte("2"))
	if err != nil {
		t.Error(err)
	}
	database.Client.Close()
}

func TestEnvUpdate(t *testing.T) {
	var adapter = InitEnv()
	err := adapter.Update("testProject", "testType", []byte("VAR_TEST=2"))
	if err != nil {
		t.Error(err)
	}
	database.Client.Close()
}

func TestEnvUpdateKey(t *testing.T) {
	var adapter = InitEnv()
	err := adapter.UpdateKey("testProject", "testType", "VAR_TEST_TWO", []byte("3"))
	if err != nil {
		t.Error(err)
	}
	database.Client.Close()
}

func TestEnvGet(t *testing.T) {
	var adapter = InitEnv()
	data, err := adapter.Get("testProject", "testType")
	if err != nil {
		t.Error(err)
	}
	if string(data) == "" {
		t.Error("Data must be VAR_TEST=2\nVAR_TEST_TWO=3")
	}
	database.Client.Close()
}

func TestEnvGetKey(t *testing.T) {
	var adapter = InitEnv()
	data, err := adapter.GetKey("testProject", "testType", "VAR_TEST_TWO")
	if err != nil {
		t.Error(err)
	}
	if string(data) != "3" {
		t.Error("Data must be 3.", string(data), "returned")
	}
	database.Client.Close()
}

func TestEnvDelete(t *testing.T) {
	var adapter = InitEnv()
	err := adapter.Delete("testProject", "testType")
	if err != nil {
		t.Error(err)
	}
	database.Client.Close()
}

func TestEnvDeleteKey(t *testing.T) {
	var adapter = InitEnv()
	err := adapter.CreateKey("testProject", "testType", "VAR_TEST_TWO", []byte("2"))
	if err != nil {
		t.Error(err)
	}

	err = adapter.DeleteKey("testProject", "testType", "VAR_TEST_TWO")
	if err != nil {
		t.Error(err)
	}
	database.Client.Close()
}
