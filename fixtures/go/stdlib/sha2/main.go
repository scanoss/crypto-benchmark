package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func main() {
	data := []byte("benchmark payload")
	sum := sha256.Sum256(data)

	h := sha256.New()
	h.Write([]byte("benchmark "))
	h.Write([]byte("payload"))
	streamed := h.Sum(nil)

	sum512 := sha512.Sum512(data)
	fmt.Println("ok", hex.EncodeToString(sum[:]) == hex.EncodeToString(streamed), len(sum512))
}
