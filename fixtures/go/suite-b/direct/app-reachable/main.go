package main

import (
	"fmt"

	"example.com/cbench/direct/liba"
)

func main() {
	fmt.Printf("%x\n", run([]byte("hello")))
}

func run(data []byte) []byte {
	return liba.Fingerprint(data)
}
