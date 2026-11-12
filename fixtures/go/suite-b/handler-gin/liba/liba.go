package liba

import "example.com/cbench/handler-gin/libb"

func Fingerprint(data []byte) []byte {
	return libb.Digest(data)
}
