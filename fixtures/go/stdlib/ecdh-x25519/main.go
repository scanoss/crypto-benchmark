package main

import (
	"bytes"
	"crypto/ecdh"
	"crypto/hkdf"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func derive(secret []byte) []byte {
	key, err := hkdf.Key(sha256.New, secret, nil, "handshake data", 32)
	if err != nil {
		panic(err)
	}
	return key
}

func main() {
	curve := ecdh.X25519()
	alice, err := curve.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	bob, err := curve.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	aliceSecret, _ := alice.ECDH(bob.PublicKey())
	bobSecret, _ := bob.ECDH(alice.PublicKey())
	fmt.Println("ok", bytes.Equal(derive(aliceSecret), derive(bobSecret)))
}
