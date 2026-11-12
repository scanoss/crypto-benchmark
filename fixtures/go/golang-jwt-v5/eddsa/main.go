package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.MapClaims{"sub": "bench"})
	signed, err := token.SignedString(privateKey)
	if err != nil {
		panic(err)
	}

	parsed, err := jwt.Parse(signed, func(t *jwt.Token) (any, error) {
		return publicKey, nil
	}, jwt.WithValidMethods([]string{"EdDSA"}))
	if err != nil {
		panic(err)
	}
	fmt.Println("ok", parsed.Valid)
}
