package libb

import "crypto/sha256"

func Digest(data []byte) []byte {
	sum := sha256.Sum256(data)
	return sum[:]
}
