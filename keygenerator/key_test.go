package keygenerator

import "testing"

func TestDummyKeyCreate(t *testing.T) {
	storage := make(map[string]bool)
	for i := 0; i < 1000; i++ {
		v, err := Key.Create()
		if _, ok := storage[v]; ok {
			t.Error("Not unique")
		} else {
			storage[v] = true
		}
		if err != nil {
			t.Error("Should be nilled")
		}
	}
}
