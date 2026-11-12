package liba

import "example.com/cbench/dispatch/libb"

type Hasher interface {
	Sum(data []byte) []byte
}

type standard struct{}

func (standard) Sum(data []byte) []byte {
	return libb.Digest(data)
}

type plain struct{}

func (plain) Sum(data []byte) []byte {
	return append([]byte(nil), data...)
}

func Standard() Hasher {
	return standard{}
}

func Plain() Hasher {
	return plain{}
}
