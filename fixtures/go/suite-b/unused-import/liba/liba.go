package liba

import "example.com/cbench/unused-import/libb"

const Version = "1.0.0"

func Fingerprint(data []byte) []byte {
	return libb.Digest(data)
}
