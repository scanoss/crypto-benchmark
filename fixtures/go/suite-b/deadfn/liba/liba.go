package liba

import "example.com/cbench/deadfn/libb"

func Sign(data []byte) []byte {
	return libb.Sign(data)
}

func Checksum(data []byte) byte {
	return libb.Checksum(data)
}
