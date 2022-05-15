package keygenerator

import "testing"

func TestDummyKeyCreate(t *testing.T) {
	key := GetKeyBuilder()
	key.Create()
	if key.Create() != "test" {
		t.Error("Sjould return test")
	}
}
