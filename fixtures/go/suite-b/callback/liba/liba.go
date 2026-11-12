package liba

import "example.com/cbench/callback/libb"

func Fingerprint(data []byte) []byte {
	return libb.Digest(data)
}
