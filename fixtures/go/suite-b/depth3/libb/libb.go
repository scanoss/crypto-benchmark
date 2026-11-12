package libb

import "example.com/cbench/depth3/libc"

func Seal(data []byte) []byte {
	return libc.Mac(data)
}
