package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	accountKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	manager := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example.com", "www.example.com"),
		Cache:      autocert.DirCache(filepath.Join(os.TempDir(), "bench-autocert")),
		Client:     &acme.Client{Key: accountKey, DirectoryURL: autocert.DefaultACMEDirectory},
	}
	server := &http.Server{
		Addr:      ":https",
		TLSConfig: manager.TLSConfig(),
	}
	fmt.Println("ok", server.TLSConfig.NextProtos)
}
