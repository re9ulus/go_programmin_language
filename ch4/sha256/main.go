package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	// task 4.1 count number of different bits
	n := 0
	for i := 0; i < len(c1); i++ {
		n += countBits(c1[i] ^ c2[i])
	}

	fmt.Printf("different bits: %d\n", n)
}

func countBits(x byte) int {
	n := 0
	for i := 0; i < 8; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}
