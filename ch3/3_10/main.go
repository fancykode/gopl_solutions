/* Exercise 3.10: Write a non-recursive version of comma, using bytes.Buffer
   instead of string concatenation.
*/
package main

import (
	"bytes"
	"fmt"
)

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	commaIndx := n % 3
	if commaIndx == 0 {
		commaIndx = 3
	}
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		if i == commaIndx {
			buf.WriteByte(',')
			commaIndx += 3
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

func main() {
	a := []string{"1",
		"12",
		"123",
		"1234",
		"12345",
		"123456",
		"1234567",
		"12345678",
		"123456789",
		"1234567890",
		"12345678901",
	}
	for i := 0; i < len(a); i++ {
		fmt.Printf("%s %s\n", a[i], comma(a[i]))
	}
}
