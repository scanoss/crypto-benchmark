package main

import (
	"fmt"

	"example.com/cbench/wrapper/liba"
)

func main() {
	fmt.Println("hello")
}

func run(data []byte) []byte {
	return liba.Protect(data)
}
