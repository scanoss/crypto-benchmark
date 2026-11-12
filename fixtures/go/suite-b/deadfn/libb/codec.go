package libb

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
)

var key = []byte("cbench-suite-b-key")

func Sign(data []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}

func Checksum(data []byte) byte {
	var c byte
	for _, b := range data {
		c ^= b
	}
	return c
}

func LegacyDigest(data []byte) []byte {
	sum := sha1.Sum(data)
	return sum[:]
}
