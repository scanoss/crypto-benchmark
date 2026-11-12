package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func main() {
	_, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	signer, err := ssh.NewSignerFromKey(privateKey)
	if err != nil {
		panic(err)
	}
	authorized := ssh.MarshalAuthorizedKey(signer.PublicKey())
	fingerprint := ssh.FingerprintSHA256(signer.PublicKey())

	data := []byte("benchmark ssh signed data")
	signature, err := signer.Sign(rand.Reader, data)
	if err != nil {
		panic(err)
	}
	err = signer.PublicKey().Verify(data, signature)
	fmt.Println("ok", err == nil, len(authorized) > 0, fingerprint[:7])
}
