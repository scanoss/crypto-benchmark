package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

func main() {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		panic(err)
	}
	plaintext := []byte("benchmark payload")
	ciphertext := gcm.Seal(nil, nonce, plaintext, []byte("header"))
	recovered, err := gcm.Open(nil, nonce, ciphertext, []byte("header"))
	if err != nil {
		panic(err)
	}
	fmt.Println("ok", bytes.Equal(recovered, plaintext))
}
