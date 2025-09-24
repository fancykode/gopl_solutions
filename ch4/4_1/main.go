/* Exercise 4.1: Write a function that counts the number of bits that are different
   in two SHA256 hashes. (See PopCount from Section 2.6.2.)
*/
package main

import (
	"crypto/sha256"
	"fmt"
)

func PopCount(n byte) int {
	count := 0
	for n != 0 {
		n = n & (n - 1)
		count++
	}
	return count
}

func CountDiffBits(h1, h2 [32]byte) int {
	diffBits := 0
	for i := 0; i < len(h1); i++ {
		diffBits += PopCount(h1[i] ^ h2[i])
	}
	return diffBits
}

func main() {
	c1 := sha256.Sum256([]byte("hello world"))
	c2 := sha256.Sum256([]byte("Hello World"))

	fmt.Printf("%x\n", c1)
	fmt.Printf("%x\n", c2)
	fmt.Printf("Number of bits that are different: %d\n", CountDiffBits(c1, c2))
}
