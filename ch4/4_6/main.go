/* Exercise 4.6: Write an in-place function that squashes each run of adjacent
   Unicode spaces (see unicode.IsSpace) in a UTF-8-encoded []byte slice into a single ASCII space.
*/
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squashSpaces(s []byte) []byte {
	result := s[:0]
	var prevRune rune = utf8.RuneError
	for i := 0; i < len(s); {
		r, sz := utf8.DecodeRune(s[i:])
		if !unicode.IsSpace(r) {
			result = append(result, s[i:i+sz]...)
		} else if unicode.IsSpace(r) && !unicode.IsSpace(prevRune) {
			result = append(result, ' ')
		}
		prevRune = r
		i += sz
	}
	return result
}

func main() {
	s := "a\n\n\tbbc\ne  d \n \n \tf\u0085\u0085g"
	result := squashSpaces([]byte(s))
	fmt.Println(string(result))
}
