package keeper

import "errors"

var Keep = NewKeeper()

type Keeper interface {
	Get(key string) (string, error)
	Set(key string, message string) error
	Delete(key string)
}

type DummyKepper struct {
	mem map[string]string
}

func (d DummyKepper) Get(key string) (string, error) {
	v, ok := d.mem[key]
	if !ok {
		return "", errors.New("message not found")
	}
	d.Delete(key)
	return v, nil
}

func (d DummyKepper) Set(key, message string) error {
	if _, ok := d.mem[key]; !ok {
		d.mem[key] = message
		return nil
	}
	return errors.New("ALready exists")
}

func (d DummyKepper) Delete(key string) {
	delete(d.mem, key)
}

func NewKeeper() Keeper {
	return DummyKepper{make(map[string]string)}
}
