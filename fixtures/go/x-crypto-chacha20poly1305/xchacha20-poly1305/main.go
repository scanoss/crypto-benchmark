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
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		panic(err)
	}
	nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+64)
	if _, err := rand.Read(nonce); err != nil {
		panic(err)
	}
	sealed := aead.Seal(nonce, nonce, []byte("benchmark payload"), []byte("v1"))
	recovered, err := aead.Open(nil, sealed[:aead.NonceSize()], sealed[aead.NonceSize():], []byte("v1"))
	if err != nil {
		panic(err)
	}
	fmt.Println("ok", bytes.Equal(recovered, []byte("benchmark payload")))
}
