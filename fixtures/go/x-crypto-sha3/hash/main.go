package main

import (
	"fmt"

	"golang.org/x/crypto/sha3"
)

func main() {
	data := []byte("benchmark payload")
	h := sha3.New256()
	h.Write(data)
	digest := h.Sum(nil)

	sum512 := sha3.Sum512(data)

	keccak := sha3.NewLegacyKeccak256()
	keccak.Write(data)
	address := keccak.Sum(nil)[12:]

	out := make([]byte, 64)
	sha3.ShakeSum256(out, data)
	fmt.Println("ok", len(digest), len(sum512), len(address), len(out))
}
