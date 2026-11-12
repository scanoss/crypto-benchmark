package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	key := make([]byte, 32)
	rand.Read(key)
	message := []byte("benchmark message")

	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	tag := mac.Sum(nil)

	check := hmac.New(sha256.New, key)
	check.Write(message)
	fmt.Println("ok", hmac.Equal(tag, check.Sum(nil)))
}
