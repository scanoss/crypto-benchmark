package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/ripemd160"
)

func main() {
	publicKey, _ := hex.DecodeString("0250863ad64a87ae8a2fe83c1af1a8403cb53f53e486d8511dad8a04887e5b2352")
	sha := sha256.Sum256(publicKey)
	h := ripemd160.New()
	h.Write(sha[:])
	fmt.Println("ok", hex.EncodeToString(h.Sum(nil)))
}
