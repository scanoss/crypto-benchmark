package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

func main() {
	password := []byte("correct horse battery staple")
	salt := make([]byte, 16)
	rand.Read(salt)
	key := pbkdf2.Key(password, salt, 600000, 32, sha256.New)
	fmt.Println("ok", len(key))
}
