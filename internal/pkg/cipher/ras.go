package cipher

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/base64"
)

type RSAEncrypter struct {
	key *rsa.PublicKey
}

type RSADecrypter struct {
	key *rsa.PrivateKey
}

func NewRSAEncrypter(key *rsa.PublicKey) *RSAEncrypter {
	return &RSAEncrypter{key: key}
}

func (e *RSAEncrypter) Encrypt(plaintext string) (string, error) {
	encrptedBytes, err := rsa.EncryptOAEP(
		sha512.New(),
		rand.Reader,
		e.key,
		[]byte(plaintext),
		nil,
	)
	if err != nil {
		return "", nil
	}

	return base64.StdEncoding.EncodeToString(encrptedBytes), nil
}

func NewRSADecrypter(key *rsa.PrivateKey) *RSADecrypter {
	return &RSADecrypter{key: key}
}

func (d *RSADecrypter) Decrypt(ciphertext []byte) ([]byte, error) {
	return nil, nil
}
