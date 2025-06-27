/*
Exercise 2.2: Write a general-purpose unit-conversion program analogous to cf that reads
numbers from its command-line arguments or from the standard input if there are no
arguments, and converts each number into units like temperature in Celsius and Fahrenheit,
length in feet and meters, weight in pounds and kilograms, and the like.
*/
package main

import (
	"2_2/lenconv"
	"2_2/tempconv"
	"2_2/weightconv"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convert(arg string) {
	v, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "2_2: %v\n", err)
		os.Exit(1)
	}
	f := tempconv.Fahrenheit(v)
	c := tempconv.Celsius(v)
	feet := lenconv.Foot(v)
	meters := lenconv.Meter(v)
	pounds := weightconv.Pound(v)
	kilograms := weightconv.Kilogram(v)

	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))

	fmt.Printf("%s = %s, %s = %s\n",
		pounds, weightconv.PndToKg(pounds),
		kilograms, weightconv.KgToPnd(kilograms))

	fmt.Printf("%s = %s, %s = %s\n",
		feet, lenconv.FtToMt(feet),
		meters, lenconv.MtToFt(meters))
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			convert(arg)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			convert(scanner.Text())
		}
	}
}
