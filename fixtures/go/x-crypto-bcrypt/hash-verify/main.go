package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := []byte("correct horse battery staple")
	hash, err := bcrypt.GenerateFromPassword(password, 12)
	if err != nil {
		panic(err)
	}
	err = bcrypt.CompareHashAndPassword(hash, password)
	fmt.Println("ok", err == nil)
}
