package main

import (
	"bytes"
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/blake2b"
)

func main() {
	key := make([]byte, 32)
	rand.Read(key)
	mac, err := blake2b.New256(key)
	if err != nil {
		panic(err)
	}
	mac.Write([]byte("benchmark message"))
	tag := mac.Sum(nil)

	check, _ := blake2b.New256(key)
	check.Write([]byte("benchmark message"))
	fmt.Println("ok", bytes.Equal(tag, check.Sum(nil)))
}
