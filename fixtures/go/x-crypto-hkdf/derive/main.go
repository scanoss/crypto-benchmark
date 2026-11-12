package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"

	"golang.org/x/crypto/hkdf"
)

func main() {
	secret := make([]byte, 32)
	rand.Read(secret)
	salt := make([]byte, sha256.Size)
	rand.Read(salt)
	reader := hkdf.New(sha256.New, secret, salt, []byte("benchmark context"))
	key := make([]byte, 32)
	if _, err := io.ReadFull(reader, key); err != nil {
		panic(err)
	}
	fmt.Println("ok", len(key))
}
