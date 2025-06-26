/* Exercise 2.1: Add types, constants, and functions to tempconv for processing temperatures in
the Kelvin scale, where zero Kelvin is −273.15°C and a difference of 1K has the same magnitude as 1°C.
*/
package main

import (
	"2_1/tempconv"
	"fmt"
)

func main() {
	k0 := tempconv.Kelvin(0)
	c0 := tempconv.Celsius(0)
	c1 := tempconv.Celsius(1)
	fmt.Println(k0, tempconv.KToC(k0))
	fmt.Println(c0, tempconv.CToK(c0))
	fmt.Println(c1, tempconv.CToK(c1))
	fmt.Println(tempconv.Celsius(100), tempconv.CToK(tempconv.Celsius(100)))
	fmt.Println(tempconv.AbsoluteZeroC, tempconv.CToK(tempconv.AbsoluteZeroC), tempconv.CToF(tempconv.AbsoluteZeroC))
	fmt.Println(tempconv.AbsoluteZeroK, tempconv.KToC(tempconv.AbsoluteZeroK),
		tempconv.CToF(tempconv.KToC(tempconv.AbsoluteZeroK)))
	fmt.Println(tempconv.BoilingC, tempconv.CToK(tempconv.BoilingC), tempconv.CToF(tempconv.BoilingC))

}
