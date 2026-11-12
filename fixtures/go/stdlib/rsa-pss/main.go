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
	opts := &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash}
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, digest[:], opts)
	if err != nil {
		panic(err)
	}
	err = rsa.VerifyPSS(&privateKey.PublicKey, crypto.SHA256, digest[:], signature, opts)
	fmt.Println("ok", err == nil)
}
