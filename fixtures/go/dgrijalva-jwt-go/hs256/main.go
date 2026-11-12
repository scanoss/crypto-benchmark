package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	secret := []byte("benchmark-hmac-secret-of-32-byte")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{Subject: "bench", ExpiresAt: time.Now().Add(time.Hour).Unix()})
	signed, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}

	parsed, err := jwt.Parse(signed, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("ok", parsed.Valid)
}
