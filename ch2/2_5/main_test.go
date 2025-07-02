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

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0)
		PopCount(1)
		PopCount(124567)
		PopCount(16428)
		PopCount(964129)
		PopCount(834255)
		PopCount(27356)
		PopCount(57612)
		PopCount(184467440737095516)
	}
}
