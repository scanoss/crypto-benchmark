package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{"sub": "bench"})
	signed, err := token.SignedString(privateKey)
	if err != nil {
		panic(err)
	}

	parsed, err := jwt.Parse(signed, func(t *jwt.Token) (any, error) {
		return &privateKey.PublicKey, nil
	}, jwt.WithValidMethods([]string{"ES256"}))
	if err != nil {
		panic(err)
	}
	fmt.Println("ok", parsed.Valid)
}
