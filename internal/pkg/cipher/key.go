package cipher

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

const DefaultRasKeySize = 2024

func GenerateRsa(size int) (private, public []byte, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		return
	}
	err = privateKey.Validate()
	if err != nil {
		return
	}

	publicBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return
	}
	public = pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicBytes,
	})

	privateBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	private = pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateBytes,
	})

	return

}
