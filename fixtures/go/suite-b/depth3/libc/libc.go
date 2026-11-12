package libc

import (
	"crypto/hmac"
	"crypto/sha256"
)

var key = []byte("cbench-suite-b-key")

func Mac(data []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}
