// Exercise 1.3: Experiment to measure the difference in running time between our potentially
// inefficient versions and the one that uses strings.Join. (Section 1.6 illustrates part of the
// time package, and Section 11.4 shows how to write benchmark tests for systematic performance evaluation.)
package main

import (
	"strings"
)

func echo1(args []string) string {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}

func echo2(args []string) string {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo3(args []string) string {
	return strings.Join(args[1:], " ")
}

func main() {
	//fmt.Println(echo1(os.Args))
	//fmt.Println(echo2(os.Args))
	//fmt.Println(echo3(os.Args))
}
