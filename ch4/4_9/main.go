/* Exercise 4.9: Write a program wordfreq to report the frequency of each word in an input text
   file. Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into words
   instead of lines.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordsCount := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		wordsCount[word]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("word\tcount\n")
	for w, c := range wordsCount {
		fmt.Printf("%q\t%d\n", w, c)
	}
}
