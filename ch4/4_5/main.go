/* Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice. */
package main

import (
	"fmt"
)

func remAdjDup(strings []string) []string {
	prev := ""
	k := 0
	for i := 0; i < len(strings); i++ {
		if strings[i] != prev {
			strings[k] = strings[i]
			k++
		}
		prev = strings[i]
	}
	return strings[:k]
}

func main() {
	data := []string{"a", "a", "a", "b", "b", "b", "c", "e", "e", "d", "f", "f", "f", "f", "k"}
	fmt.Printf("%q\n", data)
	data = remAdjDup(data)
	fmt.Printf("%q\n", data)
}
