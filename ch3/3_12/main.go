/* Exercise 3.12: Write a function that reports whether two strings are anagrams of each other,
   that is, they contain the same letters in a different order.
*/
package main

import (
	"fmt"
)

func areAnagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	runeCount := make(map[rune]int)
	for _, r := range s1 {
		runeCount[r]++
	}
	for _, r := range s2 {
		if count, ok := runeCount[r]; ok {
			if count > 1 {
				runeCount[r]--
			} else {
				delete(runeCount, r)
			}
		} else {
			return false
		}
	}
	return len(runeCount) == 0
}

func main() {
	fmt.Println("angel", "glean", areAnagrams("angel", "glean"))
	fmt.Println("arc", "car", areAnagrams("arc", "car"))
	fmt.Println("bored", "robed", areAnagrams("bored", "robed"))
	fmt.Println("cat", "act", areAnagrams("cat", "act"))
	fmt.Println("peach", "cheap", areAnagrams("peach", "cheap"))
	fmt.Println("state", "taste", areAnagrams("state", "taste"))
}
