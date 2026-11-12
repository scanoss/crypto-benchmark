package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	privatePEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
	publicDER, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(privatePEM)
	loaded, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	publicKey, err := x509.ParsePKIXPublicKey(publicDER)
	if err != nil {
		panic(err)
	}
	fmt.Println("ok", loaded.Equal(privateKey), publicKey.(*rsa.PublicKey).Equal(&privateKey.PublicKey))
}
