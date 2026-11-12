package main

import (
	"fmt"

	"example.com/cbench/transitive/liba"
)

func main() {
	fmt.Println("hello")
}

func run(data []byte) []byte {
	return liba.Fingerprint(data)
}
