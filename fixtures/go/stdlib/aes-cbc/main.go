package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

func pad(data []byte) []byte {
	n := aes.BlockSize - len(data)%aes.BlockSize
	return append(data, bytes.Repeat([]byte{byte(n)}, n)...)
}

func main() {
	key := make([]byte, 32)
	rand.Read(key)
	iv := make([]byte, aes.BlockSize)
	rand.Read(iv)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	plaintext := pad([]byte("benchmark payload"))
	ciphertext := make([]byte, len(plaintext))
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(ciphertext, plaintext)

	recovered := make([]byte, len(ciphertext))
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(recovered, ciphertext)
	fmt.Println("ok", bytes.Equal(recovered, plaintext))
}
