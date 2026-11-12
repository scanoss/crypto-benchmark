package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	generated, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(generated)})
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyPEM)
	if err != nil {
		panic(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{"sub": "bench"})
	signed, err := token.SignedString(privateKey)
	if err != nil {
		panic(err)
	}

	parsed, err := jwt.Parse(signed, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return &privateKey.PublicKey, nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("ok", parsed.Valid)
}
