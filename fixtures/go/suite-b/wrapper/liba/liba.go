package liba

import "example.com/cbench/wrapper/libb"

func Protect(data []byte) []byte {
	return protect(data)
}

func protect(data []byte) []byte {
	return libb.Encode(data)
}
