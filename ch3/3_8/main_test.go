package main

import (
	"io"
	"testing"
)

func BenchmarkCalcMandelbrot64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calcMandelbrot64(io.Discard)
	}
}

func BenchmarkCalcMandelbrot128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calcMandelbrot128(io.Discard)
	}
}

func BenchmarkCalMandelbrotBigFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calcMandelbrotBigFloat(io.Discard)
	}
}

func BenchmarkCalMandelbrotBigRat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calcMandelbrotBigRat(io.Discard)
	}
}
