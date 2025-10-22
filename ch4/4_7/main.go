/* Exercise 4.7: Modify reverse to reverse the characters of a []byte slice that
   represents a UTF-8-encoded string, in place. Can you do it without allocating new memory?
*/

package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseUTF8(s []byte) {
	for i := 0; i < len(s); {
		_, sz := utf8.DecodeRune(s[i:])
		reverse(s[i : i+sz])
		i += sz
	}
	reverse(s)
}

func main() {
	s := []byte("Hello, 世界\U0001f44d\U0001F525")
	fmt.Println(string(s))
	reverseUTF8(s)
	fmt.Println(string(s))
}
