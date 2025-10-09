/* Exercise 4.4: Write a version of rotate that operates in a single pass. */
package main

import (
	"fmt"
)

func rotateLeft(slice []int, n int) []int {
	if n < 0 {
		return nil
	}
	n = n % len(slice)
	result := make([]int, len(slice))

	for i, j := n, 0; j < len(result); i, j = i+1, j+1 {
		result[j] = slice[i%len(slice)]
	}
	return result
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(a)
	fmt.Println(rotateLeft(a, 2))
}
