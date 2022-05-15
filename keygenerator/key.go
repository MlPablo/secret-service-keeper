package keygenerator

var Key_builder = GetKeyBuilder()

type KeyBuilder interface {
	Create() string
}

type DummyKey struct{}

func (k DummyKey) Create() string {
	return "test"
}

func GetKeyBuilder() KeyBuilder {
	return DummyKey{}
}
