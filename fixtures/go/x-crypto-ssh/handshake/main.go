package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
)

func main() {
	_, hostKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	hostSigner, err := ssh.NewSignerFromKey(hostKey)
	if err != nil {
		panic(err)
	}
	serverConfig := &ssh.ServerConfig{
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			if c.User() == "bench" && string(pass) == "s3cret" {
				return nil, nil
			}
			return nil, fmt.Errorf("password rejected for %q", c.User())
		},
	}
	serverConfig.AddHostKey(hostSigner)

	clientConfig := &ssh.ClientConfig{
		User:            "bench",
		Auth:            []ssh.AuthMethod{ssh.Password("s3cret")},
		HostKeyCallback: ssh.FixedHostKey(hostSigner.PublicKey()),
	}

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	go func() {
		serverSide, err := listener.Accept()
		if err != nil {
			return
		}
		conn, chans, reqs, err := ssh.NewServerConn(serverSide, serverConfig)
		if err != nil {
			return
		}
		go ssh.DiscardRequests(reqs)
		for ch := range chans {
			ch.Reject(ssh.Prohibited, "no channels")
		}
		conn.Close()
	}()
	clientSide, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		panic(err)
	}
	conn, _, _, err := ssh.NewClientConn(clientSide, listener.Addr().String(), clientConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println("ok", string(conn.ServerVersion()))
	conn.Close()
}
