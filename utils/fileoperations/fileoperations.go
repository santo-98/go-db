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

func getGCM(key []byte) cipher.AEAD {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	return gcm
}

func ReadEncryptedFile(path string) {
	key := getKeys()

	ciphertext, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	gcm := getGCM(key)

	nonce := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(plaintext))
}

func WriteEncryptedFile() {
	key := getKeys()

	gcm := getGCM(key)

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte("plaintext"), nil)

	err := os.WriteFile("test.godb", ciphertext, 0777)
	if err != nil {
		panic(err)
	}
}
