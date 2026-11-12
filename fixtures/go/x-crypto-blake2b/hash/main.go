package main

import (
	"fmt"

	"golang.org/x/crypto/blake2b"
)

func main() {
	data := []byte("benchmark payload")
	sum := blake2b.Sum256(data)

	h, err := blake2b.New512(nil)
	if err != nil {
		panic(err)
	}
	h.Write(data)
	fmt.Println("ok", len(sum), len(h.Sum(nil)))
}
