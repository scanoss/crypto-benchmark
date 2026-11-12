package main

import (
	"bytes"
	"fmt"
	"slices"
	"strings"

	"example.com/cbench/callback/liba"
)

func main() {
	names := []string{"carol", "alice", "bob"}
	slices.SortFunc(names, byFingerprint)
	fmt.Println(strings.Join(names, " "))
}

// byFingerprint orders names by their fingerprint, a stable shuffle.
func byFingerprint(a, b string) int {
	return bytes.Compare(liba.Fingerprint([]byte(a)), liba.Fingerprint([]byte(b)))
}
