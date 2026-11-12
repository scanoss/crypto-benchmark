package main

import (
	"fmt"

	"example.com/cbench/wrapper/liba"
)

func main() {
	fmt.Printf("%x\n", run([]byte("hello")))
}

func run(data []byte) []byte {
	return liba.Protect(data)
}
