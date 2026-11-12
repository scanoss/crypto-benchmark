package main

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	hash := crypto.Keccak256([]byte("benchmark message"))
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		panic(err)
	}

	recovered, err := crypto.Ecrecover(hash, signature)
	if err != nil {
		panic(err)
	}
	publicKey := crypto.FromECDSAPub(&privateKey.PublicKey)
	valid := crypto.VerifySignature(publicKey, hash, signature[:64])
	fmt.Println("ok", bytes.Equal(recovered, publicKey), valid)
}
