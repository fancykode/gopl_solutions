package main

import "testing"

func BenchmarkPopCountTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountTable(0)
		PopCountTable(1)
		PopCountTable(124567)
		PopCountTable(16428)
		PopCountTable(964129)
		PopCountTable(834255)
		PopCountTable(27356)
		PopCountTable(57612)
		PopCountTable(184467440737095516)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(0)
		PopCountShift(1)
		PopCountShift(124567)
		PopCountShift(16428)
		PopCountShift(964129)
		PopCountShift(834255)
		PopCountShift(27356)
		PopCountShift(57612)
		PopCountShift(184467440737095516)
	}
}
