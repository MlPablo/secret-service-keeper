package keeper

import (
	"testing"
)

func TestKeeperSet(t *testing.T) {
	keeper := DummyKepper{make(map[string]string)}
	key := "foo"
	text := "bar"
	keeper.Set(key, text, 0)
	if keeper.mem[key] == text {
		t.Error("Should be encrypted")
	}
	err := keeper.Set(key, text, 0)
	if err == nil {
		t.Error("Should be errored")
	}
}

func TestKeeperGet(t *testing.T) {
	keeper := DummyKepper{make(map[string]string)}
	key := "foo"
	text := "bar"
	keeper.Set(key, text, 0)
	v, _ := keeper.Get(key)
	if v != text {
		t.Error("Should be returned")
	}
	if len(keeper.mem) != 0 {
		t.Error("Should be deleted")
	}
	_, err := keeper.Get(key)
	if err == nil {
		t.Error("Should be errored")
	}

}

func TestKeeperDelete(t *testing.T) {
	keeper := DummyKepper{make(map[string]string)}
	key := "foo"
	text := "bar"
	keeper.Set(key, text, 0)
	keeper.Delete(key)
	if len(keeper.mem) != 0 {
		t.Error("Should be deleted")
	}
}
