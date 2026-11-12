package main

import (
	"bytes"
	"crypto/mlkem"
	"fmt"
)

func main() {
	decapsulationKey, err := mlkem.GenerateKey768()
	if err != nil {
		panic(err)
	}
	encapsulationKey := decapsulationKey.EncapsulationKey()
	sharedKey, ciphertext := encapsulationKey.Encapsulate()
	recovered, err := decapsulationKey.Decapsulate(ciphertext)
	if err != nil {
		panic(err)
	}
	fmt.Println("ok", bytes.Equal(sharedKey, recovered))
}
