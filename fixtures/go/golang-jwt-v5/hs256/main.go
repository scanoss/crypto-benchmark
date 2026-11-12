package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	secret := []byte("benchmark-hmac-secret-of-32-byte")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": "bench", "exp": time.Now().Add(time.Hour).Unix()})
	signed, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}

	parsed, err := jwt.Parse(signed, func(t *jwt.Token) (any, error) {
		return secret, nil
	}, jwt.WithValidMethods([]string{"HS256"}))
	if err != nil {
		panic(err)
	}
	fmt.Println("ok", parsed.Valid)
}
