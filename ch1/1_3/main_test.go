package main

import "testing"

var args = []string{"program", "Argument0", "Argument1",
	"Argument2", "Argument3", "Argument4", "Argument5",
	"Argument6", "Argument7", "Argument8", "Argument9",
	"Argument10", "Argument11", "Argument12", "Argument13",
	"Argument14", "Argument15", "Argument16", "Argument17",
	"Argument18", "Argument19", "Argument20", "Argument21",
	"Argument22", "Argument23", "Argument24", "Argument25",
	"Argument26", "Argument27", "Argument28", "Argument29",
	"Argument30", "Argument31", "Argument32", "Argument33",
	"Argument34", "Argument35", "Argument36", "Argument37",
	"Argument38", "Argument39", "Argument40", "Argument41",
}

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(args)
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2(args)
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3(args)
	}
}
