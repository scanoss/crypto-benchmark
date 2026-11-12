package libb

import "crypto/sha512"

func Digest(data []byte) []byte {
	sum := sha512.Sum512(data)
	return sum[:]
}
