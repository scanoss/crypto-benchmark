package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"crypto/rc4"
	"fmt"
)

func main() {
	key := make([]byte, 24)
	rand.Read(key)
	iv := make([]byte, des.BlockSize)
	rand.Read(iv)
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}
	plaintext := []byte("16 byte payload!")
	ciphertext := make([]byte, len(plaintext))
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(ciphertext, plaintext)

	streamKey := []byte("legacy-rc4-key")
	stream, err := rc4.NewCipher(streamKey)
	if err != nil {
		panic(err)
	}
	obfuscated := make([]byte, len(plaintext))
	stream.XORKeyStream(obfuscated, plaintext)
	fmt.Println("ok", !bytes.Equal(ciphertext, plaintext), !bytes.Equal(obfuscated, plaintext))
}
