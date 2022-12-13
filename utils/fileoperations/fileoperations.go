package fileoperations

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func getKeys() []byte {
	key, err := os.ReadFile("config/keys")
	if err != nil {
		panic(err)
	}

	return key
}

func ReadEncryptedFile() {
	key := getKeys()
	fmt.Println(string(key))
}

func WriteEncryptedFile() {
	key := getKeys()

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte("plaintext"), nil)

	err = os.WriteFile("test.godb", ciphertext, 0777)
	if err != nil {
		panic(err)
	}
}
