package libb

import "crypto/sha512"

func Encode(data []byte) []byte {
	return encode(data)
}

func encode(data []byte) []byte {
	return hashBlock(data)
}

func hashBlock(data []byte) []byte {
	sum := sha512.Sum384(data)
	return sum[:]
}
