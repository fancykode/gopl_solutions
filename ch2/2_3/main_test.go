package main

import "testing"

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(0)
		PopCountLoop(1)
		PopCountLoop(124567)
		PopCountLoop(16428)
		PopCountLoop(964129)
		PopCountLoop(834255)
		PopCountLoop(27356)
		PopCountLoop(57612)
		PopCountLoop(184467440737095516)
	}
}

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
