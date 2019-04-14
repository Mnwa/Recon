package main_test

import (
	"Recon/adapters"
	"testing"
)

func TestConfigCreate(t *testing.T) {
	var adapter = adapters.GetDefault()
	err := adapter.Create("testProject", "testType", []byte("VAR_TEST=1"))
	if err != nil {
		t.Error(err)
	}
}

func TestConfigCreateKey(t *testing.T) {
	var adapter = adapters.GetDefault()
	err := adapter.CreateKey("testProject", "testType", "VAR_TEST_TWO", []byte("2"))
	if err != nil {
		t.Error(err)
	}
}

func TestConfigUpdate(t *testing.T) {
	var adapter = adapters.GetDefault()
	err := adapter.Update("testProject", "testType", []byte("VAR_TEST=2"))
	if err != nil {
		t.Error(err)
	}
}

func TestConfigUpdateKey(t *testing.T) {
	var adapter = adapters.GetDefault()
	err := adapter.UpdateKey("testProject", "testType", "VAR_TEST_TWO", []byte("3"))
	if err != nil {
		t.Error(err)
	}
}

func TestConfigGet(t *testing.T) {
	var adapter = adapters.GetDefault()
	data, err := adapter.Get("testProject", "testType")
	if err != nil {
		t.Error(err)
	}
	if string(data) == "" {
		t.Error("Data must be VAR_TEST=2\nVAR_TEST_TWO=3")
	}
}

func TestConfigGetKey(t *testing.T) {
	var adapter = adapters.GetDefault()
	data, err := adapter.GetKey("testProject", "testType", "VAR_TEST_TWO")
	if err != nil {
		t.Error(err)
	}
	if string(data) != "3" {
		t.Error("Data must be 3.", string(data), "returned")
	}
}

func TestConfigDelete(t *testing.T) {
	var adapter = adapters.GetDefault()
	err := adapter.Delete("testProject", "testType")
	if err != nil {
		t.Error(err)
	}
}

func TestConfigDeleteKey(t *testing.T) {
	var adapter = adapters.GetDefault()
	err := adapter.CreateKey("testProject", "testType", "VAR_TEST_TWO", []byte("2"))
	if err != nil {
		t.Error(err)
	}

	err = adapter.DeleteKey("testProject", "testType", "VAR_TEST_TWO")
	if err != nil {
		t.Error(err)
	}
}
