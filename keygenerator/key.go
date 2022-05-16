package keygenerator

import "github.com/google/uuid"

var Key = GetKeyBuilder()

type KeyBuilder interface {
	Create() (string, error)
}

//type DummyKey struct{}
//
//func (k DummyKey) Create() (string, error) {
//	return "test", nil
//}

func GetKeyBuilder() KeyBuilder {
	return UUIDKeyGenerator{}
}

type UUIDKeyGenerator struct{}

func (k UUIDKeyGenerator) Create() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}
