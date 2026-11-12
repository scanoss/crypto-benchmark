package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	digest := sha256.Sum256([]byte("benchmark message"))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, digest[:])
	if err != nil {
		panic(err)
	}
	err = rsa.VerifyPKCS1v15(&privateKey.PublicKey, crypto.SHA256, digest[:], signature)
	fmt.Println("ok", err == nil)
}
