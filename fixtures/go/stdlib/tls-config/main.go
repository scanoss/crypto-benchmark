package main

import (
	"crypto/tls"
	"fmt"
)

func main() {
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: "example.com",
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		},
	}
	fmt.Println("ok", config.MinVersion == tls.VersionTLS12)
}
