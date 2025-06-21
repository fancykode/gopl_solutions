// Exercise 1.4: Modify dup2 to print the names of all files in
// which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/* Converts set of file names to a sorted slice */
func getFileNames(fileNames map[string]struct{}) []string {
	result := make([]string, 0, len(fileNames))
	for fname := range fileNames {
		result = append(result, fname)
	}
	sort.Sort(sort.StringSlice(result))
	return result
}

func main() {
	// key: line, value: number of this lines in all files
	counts := make(map[string]int)
	// key: line, value set of file names (as map[string]struct{}) that contain this line
	lineFileNames := make(map[string]map[string]struct{})

	files := os.Args[1:]
	if len(files) != 0 {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2mod: %v\n", err)
				continue
			}

			input := bufio.NewScanner(f)
			for input.Scan() {
				counts[input.Text()]++

				if _, ok := lineFileNames[input.Text()]; !ok {
					lineFileNames[input.Text()] = make(map[string]struct{})
				}
				lineFileNames[input.Text()][f.Name()] = struct{}{}

			}
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%q\t%v\n", n, line, getFileNames(lineFileNames[line]))
		}
	}
}
