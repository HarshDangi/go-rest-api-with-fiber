package database

import "testing"

func TestDBConnectionAndInitialization(t *testing.T) {
	err := Connect()

	if err != nil {
		t.Errorf("Couldn't establish a connection.\n Reason: %s", err.Error())
	}
}
