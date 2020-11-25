package main

import (
	"testing"
	"reflect"
)

func isType(value interface{}, expectedType interface{}) (bool) {
	return reflect.TypeOf(value).String() == expectedType
}

func TestGetConfigValues(t *testing.T) {
	duration, interval, logpath, hostIP := getConfigValues()

	testTables := []struct {
		variableName string
		value interface{}
		expectedType interface{}
	}{
		{ "duration", duration, "int" },
		{ "interval", interval, "int" },
		{ "logpath", logpath, "string" },
		{ "hostIP", hostIP, "string" },
	}

	for _, table := range testTables {
		isExpectedType := isType(table.value, table.expectedType)
		if !isExpectedType {
			t.Errorf("Expect %s with value %s to be %s", table.variableName, table.value, table.expectedType)
		}
	}
}

func TestPingHost(t *testing.T) {
	packetLoss := pingHost("127.0.0.1")
	if packetLoss > 0 {
		t.Errorf("Expect packetLoss to be 0, got %f", packetLoss)
	}
}