package main

import (
	"fmt"

	"example.com/cbench/depth3/liba"
)

func main() {
	fmt.Printf("%x\n", run([]byte("hello")))
}

func run(data []byte) []byte {
	return liba.Seal(data)
}
