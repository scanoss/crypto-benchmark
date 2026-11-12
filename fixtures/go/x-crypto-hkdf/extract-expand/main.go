package main

import (
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"io"

	"golang.org/x/crypto/hkdf"
)

func main() {
	secret := make([]byte, 64)
	rand.Read(secret)
	prk := hkdf.Extract(sha512.New, secret, []byte("benchmark salt"))
	encKey := make([]byte, 32)
	io.ReadFull(hkdf.Expand(sha512.New, prk, []byte("encryption")), encKey)
	macKey := make([]byte, 32)
	io.ReadFull(hkdf.Expand(sha512.New, prk, []byte("authentication")), macKey)
	fmt.Println("ok", len(encKey), len(macKey))
}
