package liba

import "example.com/cbench/handler-nethttp/libb"

func Fingerprint(data []byte) []byte {
	return libb.Digest(data)
}
