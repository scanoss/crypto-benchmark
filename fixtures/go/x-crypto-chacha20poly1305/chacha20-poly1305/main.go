package main

import (
	"bytes"
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/chacha20poly1305"
)

func main() {
	key := make([]byte, chacha20poly1305.KeySize)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	aead, err := chacha20poly1305.New(key)
	if err != nil {
		panic(err)
	}
	nonce := make([]byte, chacha20poly1305.NonceSize)
	if _, err := rand.Read(nonce); err != nil {
		panic(err)
	}
	plaintext := []byte("benchmark payload")
	ciphertext := aead.Seal(nil, nonce, plaintext, nil)
	recovered, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("ok", bytes.Equal(recovered, plaintext))
}
