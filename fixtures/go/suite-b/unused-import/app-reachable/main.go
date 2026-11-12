package main

import (
	"fmt"

	"example.com/cbench/unused-import/liba"
)

func main() {
	fmt.Printf("%s %x\n", liba.Version, liba.Fingerprint([]byte("hello")))
}
