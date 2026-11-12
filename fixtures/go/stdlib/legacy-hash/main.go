package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	content := []byte("legacy artifact contents")
	etag := md5.Sum(content)
	gitID := sha1.Sum(append([]byte(fmt.Sprintf("blob %d\x00", len(content))), content...))
	fmt.Printf("ok %x %x\n", etag, gitID)
}
