package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

func main() {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	keyring := agent.NewKeyring()
	if err := keyring.Add(agent.AddedKey{PrivateKey: privateKey, Comment: "bench"}); err != nil {
		panic(err)
	}

	agentSide, clientSide := net.Pipe()
	go agent.ServeAgent(keyring, agentSide)
	client := agent.NewClient(clientSide)
	keys, err := client.List()
	if err != nil {
		panic(err)
	}

	data := []byte("benchmark agent signed data")
	signature, err := client.Sign(keys[0], data)
	if err != nil {
		panic(err)
	}
	sshPublicKey, err := ssh.NewPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	err = sshPublicKey.Verify(data, signature)
	fmt.Println("ok", err == nil, keys[0].Comment)
}
