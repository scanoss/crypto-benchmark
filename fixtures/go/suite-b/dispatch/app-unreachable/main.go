package main

import (
	"fmt"

	"example.com/cbench/dispatch/liba"
)

func main() {
	h := liba.Plain()
	fmt.Printf("%x\n", h.Sum([]byte("hello")))
}
