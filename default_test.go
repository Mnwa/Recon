package main_test

import (
	"Recon/adapters"
	"Recon/database"
	"github.com/prologic/bitcask"
	"log"
	"testing"
)

func InitDefault() *adapters.Default {
	var err error
	database.Client, err = bitcask.Open("/tmp/test_recon")
	if err != nil {
		log.Fatal(err)
	}

	return adapters.GetDefault()
}

func TestConfigCreate(t *testing.T) {
	var adapter = InitDefault()
	err := adapter.Create("testProject", "testType", []byte("VAR_TEST=1"))
	if err != nil {
		t.Error(err)
	}
	database.Client.Close()
}

func TestConfigCreateKey(t *testing.T) {
	var adapter = InitDefault()
	err := adapter.CreateKey("testProject", "testType", "VAR_TEST_TWO", []byte("2"))
	if err != nil {
		t.Error(err)
	}
	database.Client.Close()
}

func TestConfigUpdate(t *testing.T) {
	var adapter = InitDefault()
	err := adapter.Update("testProject", "testType", []byte("VAR_TEST=2"))
	if err != nil {
		t.Error(err)
	}
	database.Client.Close()
}

func TestConfigUpdateKey(t *testing.T) {
	var adapter = InitDefault()
	err := adapter.UpdateKey("testProject", "testType", "VAR_TEST_TWO", []byte("3"))
	if err != nil {
		t.Error(err)
	}
	database.Client.Close()
}

func TestConfigGet(t *testing.T) {
	var adapter = InitDefault()
	data, err := adapter.Get("testProject", "testType")
	if err != nil {
		t.Error(err)
	}
	if string(data) == "" {
		t.Error("Data must be VAR_TEST=2\nVAR_TEST_TWO=3")
	}
	database.Client.Close()
}

func TestConfigGetKey(t *testing.T) {
	var adapter = InitDefault()
	data, err := adapter.GetKey("testProject", "testType", "VAR_TEST_TWO")
	if err != nil {
		t.Error(err)
	}
	if string(data) != "3" {
		t.Error("Data must be 3.", string(data), "returned")
	}
	database.Client.Close()
}

func TestConfigDelete(t *testing.T) {
	var adapter = InitDefault()
	err := adapter.Delete("testProject", "testType")
	if err != nil {
		t.Error(err)
	}
	database.Client.Close()
}

func TestConfigDeleteKey(t *testing.T) {
	var adapter = InitDefault()
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
