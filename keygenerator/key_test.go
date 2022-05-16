package keygenerator

import "testing"

func TestDummyKeyCreate(t *testing.T) {
	if Key.Create() != "test" {
		t.Error("Should return test")
	}
}
