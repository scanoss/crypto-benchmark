package main

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	keyHex := hex.EncodeToString(crypto.FromECDSA(privateKey))
	loaded, err := crypto.HexToECDSA(keyHex)
	if err != nil {
		panic(err)
	}
	address := crypto.PubkeyToAddress(loaded.PublicKey)
	fmt.Println("ok", address == crypto.PubkeyToAddress(privateKey.PublicKey), address.Hex()[:6])
}
