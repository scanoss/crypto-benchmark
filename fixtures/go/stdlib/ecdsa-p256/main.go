package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256([]byte("benchmark message"))
	signature, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}
	valid := ecdsa.VerifyASN1(&privateKey.PublicKey, hash[:], signature)
	fmt.Println("ok", valid)
}
