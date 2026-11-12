package main

import (
	"fmt"

	"example.com/cbench/dispatch/liba"
)

func main() {
	h := liba.Standard()
	fmt.Printf("%x\n", h.Sum([]byte("hello")))
}
