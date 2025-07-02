/* Exercise 2.4: Write a version of PopCount that counts bits by shifting its argument through 64
bit position s, testing the rightmost bit each time. Compare its performance to the table-lookup version.
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

// PopCountShift returns the population count (number of set bits) of x.
func PopCountShift(x uint64) int {
	var count int
	for x != 0 {
		count += int(x & 1)
		x >>= 1
	}
	return count
}
