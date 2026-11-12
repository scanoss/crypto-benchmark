package liba

import "crypto/sha256"

func Fingerprint(data []byte) []byte {
	sum := sha256.Sum256(data)
	return sum[:]
}
