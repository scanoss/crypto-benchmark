package main

import (
	"fmt"

	"example.com/cbench/depth3/liba"
)

func main() {
	fmt.Println("hello")
}

func run(data []byte) []byte {
	return liba.Seal(data)
}
