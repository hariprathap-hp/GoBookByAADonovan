package main

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	PopCount(63)
}

func BenchmarkBitpopcount(b *testing.B) {
	BitPopCount(63)
}

func BenchmarkRightShiftPopcount(b *testing.B) {
	RightShiftPopCount(63)
}

func BenchmarkClear1(b *testing.B) {
	PopCountClear1(63)
}
