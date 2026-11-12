package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 3072)
	if err != nil {
		panic(err)
	}
	passphrase := []byte("s3cret-passphrase")
	block, err := ssh.MarshalPrivateKeyWithPassphrase(privateKey, "bench@example", passphrase)
	if err != nil {
		panic(err)
	}
	path := filepath.Join(os.TempDir(), "bench_id_rsa")
	defer os.Remove(path)
	if err := os.WriteFile(path, pem.EncodeToMemory(block), 0o600); err != nil {
		panic(err)
	}

	pemBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	signer, err := ssh.ParsePrivateKeyWithPassphrase(pemBytes, passphrase)
	if err != nil {
		panic(err)
	}
	fmt.Println("ok", signer.PublicKey().Type())
}
