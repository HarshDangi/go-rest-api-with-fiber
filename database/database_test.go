package database

import "testing"

func TestDBConnectionAndInitialization(t *testing.T) {
	err := Connect("testing_products")

	if err != nil {
		t.Errorf("Couldn't establish a connection.\n Reason: %s", err.Error())
	}
}
