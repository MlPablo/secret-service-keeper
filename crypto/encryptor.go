package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

var SecretKey = []byte("secretkey1234567")

func Encrypt(msg string) (string, error) {
	c, err := aes.NewCipher(SecretKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	return string(gcm.Seal(nonce, nonce, []byte(msg), nil)), nil
}

func Decrypt(msg string) (string, error) {
	c, err := aes.NewCipher(SecretKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(msg) < nonceSize {
		return "", err
	}
	//msg = []byte(msg)
	nonce, mmsg := []byte(msg[:nonceSize]), []byte(msg[nonceSize:])
	plaintext, err := gcm.Open(nil, nonce, mmsg, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
