package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println(hammingdistance(sha256.Sum256([]byte("X")), sha256.Sum256([]byte("A"))))
}

func hammingdistance(c1 [32]byte, c2 [32]byte) int {
	num := 0
	for i := 0; i < 32; i++ {
		for j := 0; j < 8; j++ {
			num += int((c1[i]%2 + c2[i]) % 2)
			c1[i] = c1[i] >> 1
			c2[i] = c2[i] >> 1
		}
		num += int((c1[i]%2 + c2[i]) % 2)
	}
	return num
}
