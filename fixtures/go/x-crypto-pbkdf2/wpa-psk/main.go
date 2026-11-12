package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

func main() {
	passphrase := []byte("wireless passphrase")
	ssid := []byte("HomeNetwork")
	psk := pbkdf2.Key(passphrase, ssid, 4096, 32, sha1.New)
	fmt.Println("ok", hex.EncodeToString(psk)[:8])
}
