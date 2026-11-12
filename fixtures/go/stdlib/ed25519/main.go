package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
)

func main() {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	message := []byte("benchmark message")
	signature := ed25519.Sign(privateKey, message)
	fmt.Println("ok", ed25519.Verify(publicKey, message, signature))
}
