package main

import (
	"fmt"

	"example.com/cbench/transitive/liba"
)

func main() {
	fmt.Printf("%x\n", run([]byte("hello")))
}

func run(data []byte) []byte {
	return liba.Fingerprint(data)
}
