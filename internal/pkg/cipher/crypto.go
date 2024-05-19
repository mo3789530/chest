package cipher

type Encrypter interface {
	Encrypt(text string) (string, error)
}

type Decrypter interface {
	Decrypt(encryptedText string) (string, error)
}

type EncryptDecrypter interface {
	Encrypter
	Decrypter
}
