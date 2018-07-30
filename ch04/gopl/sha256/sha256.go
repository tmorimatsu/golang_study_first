package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	c1 := convertToSha256("x") // sha256.Sum256([]byte("x"))
	c2 := convertToSha256("X") // sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

func convertToSha256(s string) [32]byte {
	return sha256.Sum256([]byte(s))
}
