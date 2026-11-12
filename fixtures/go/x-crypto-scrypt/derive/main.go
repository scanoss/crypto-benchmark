package main

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

func main() {
	password := []byte("correct horse battery staple")
	salt := make([]byte, 16)
	rand.Read(salt)
	key, err := scrypt.Key(password, salt, 1<<15, 8, 1, 32)
	if err != nil {
		panic(err)
	}
	fmt.Println("ok", len(key))
}
