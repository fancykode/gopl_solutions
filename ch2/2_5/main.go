/* Exercise 2.5: The expression x&(x-1) clears the rightmost non-zero bit of x.
   Write a version of PopCount that counts bits by using this fact, and assess its performance.
*/
package main

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountTable returns the population count (number of set bits) of x.
func PopCountTable(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount(x uint64) int {
	var count int
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}
