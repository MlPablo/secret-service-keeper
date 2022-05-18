package keeper

import (
	"errors"
	"secretservice/crypto"
)

type Keeper interface {
	Get(key string) (string, error)
	Set(key string, message string, ttl int) error
	Delete(key string) error
}

type DummyKepper struct {
	mem map[string]string
}

func (d DummyKepper) Get(key string) (string, error) {
	v, ok := d.mem[key]
	if !ok {
		return "", errors.New("message not found")
	}
	encrypted, err := crypto.Decrypt(v)
	if err != nil {
		return "", err
	}
	d.Delete(key)
	return encrypted, nil
}

func (d DummyKepper) Set(key, message string, ttl int) error {
	encryptedMessage, err := crypto.Encrypt(message)
	if err != nil {
		return err
	}
	if _, ok := d.mem[key]; !ok {
		d.mem[key] = encryptedMessage
		return nil
	}
	return errors.New("ALready exists")
}

func (d DummyKepper) Delete(key string) error {
	delete(d.mem, key)
	return nil
}

func NewKeeper() Keeper {
	return DummyKepper{make(map[string]string)}
}
