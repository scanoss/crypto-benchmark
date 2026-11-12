package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	sessionID := make([]byte, 32)
	if _, err := rand.Read(sessionID); err != nil {
		panic(err)
	}
	pin, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		panic(err)
	}
	token := rand.Text()
	fmt.Println("ok", len(sessionID), pin.Sign() >= 0, len(token))
}
