package liba

import "example.com/cbench/transitive/libb"

func Fingerprint(data []byte) []byte {
	return libb.Digest(data)
}
