package main

import (
	"fmt"

	"example.com/cbench/deadfn/liba"
)

func main() {
	fmt.Printf("%x\n", liba.Sign([]byte("hello")))
}
