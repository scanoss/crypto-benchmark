package main

import (
	"crypto/sha3"
	"fmt"
)

func main() {
	data := []byte("benchmark payload")
	sum := sha3.Sum256(data)

	h := sha3.New512()
	h.Write(data)
	sum512 := h.Sum(nil)

	xof := sha3.SumSHAKE256(data, 64)
	fmt.Println("ok", len(sum), len(sum512), len(xof))
}
