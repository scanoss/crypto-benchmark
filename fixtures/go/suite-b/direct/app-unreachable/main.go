package main

import (
	"fmt"

	"example.com/cbench/direct/liba"
)

func main() {
	fmt.Println("hello")
}

func run(data []byte) []byte {
	return liba.Fingerprint(data)
}
