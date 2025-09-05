/* Exercise 3.11: Enhance comma so that it deals correctly with floating-point numbers
   and an optional sign.
*/
package main

import (
	"bytes"
	"fmt"
	"strings"
)

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var sign string
	if s[0] == '+' || s[0] == '-' {
		sign = string(s[0])
		s = s[1:]
	}
	dotIndx := strings.Index(s, ".")
	var dotPart string
	if dotIndx != -1 {
		dotPart = s[dotIndx:]
		s = s[:dotIndx]
	}
	return sign + commaInt(s) + dotPart
}

func commaInt(s string) string {
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
	a = []string{
		"0.01",
		"-0.01",
		"+0.01",
		"1",
		"+1",
		"-1",
		"+192345678",
		"-192345678",
		"192345678",
		"-1.001",
		"+1.001",
		"1.001",
		"1234.023",
		"+1234.023",
		"-1234.023",
		"12.12332432",
		"123.12334345",
		"1234.34534234",
		"12345.234123213",
		"123456.2342312",
		"1234567.5678567",
		"12345678.54654734",
		"123456789.501010401",
		"1234567890.40503002",
		"12345678901.3223423",
	}
	for i := 0; i < len(a); i++ {
		fmt.Printf("%s %s\n", a[i], comma(a[i]))
	}
}
