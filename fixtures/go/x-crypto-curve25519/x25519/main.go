package main

import (
	"bytes"
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/curve25519"
)

func main() {
	alicePrivate := make([]byte, curve25519.ScalarSize)
	rand.Read(alicePrivate)
	bobPrivate := make([]byte, curve25519.ScalarSize)
	rand.Read(bobPrivate)

	alicePublic, err := curve25519.X25519(alicePrivate, curve25519.Basepoint)
	if err != nil {
		panic(err)
	}
	bobPublic, err := curve25519.X25519(bobPrivate, curve25519.Basepoint)
	if err != nil {
		panic(err)
	}
	aliceShared, _ := curve25519.X25519(alicePrivate, bobPublic)
	bobShared, _ := curve25519.X25519(bobPrivate, alicePublic)
	fmt.Println("ok", bytes.Equal(aliceShared, bobShared))
}
