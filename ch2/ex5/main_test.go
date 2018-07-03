package main

import (
	"testing"

	"./popcount"
)

// TODO: testをそれぞれのexに対してきちんと書く

// default
func BenchmarkPopCount(b *testing.B) {
	n := 114514
	for i := 0; i < b.N; i++ {
		popcount.PopCount(uint64(n))
	}
}

// ch2/ex3
func BenchmarkPopCountLoop(b *testing.B) {
	n := 114514
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(uint64(n))
	}
}

// ch2/ex4
func BenchmarkPopCountBitShift64(b *testing.B) {
	n := 114514
	for i := 0; i < b.N; i++ {
		popcount.PopCountBitShift64(uint64(n))
	}
}

// ch2/ex3
func BenchmarkPopCountClearLowestBit(b *testing.B) {
	n := 114514
	for i := 0; i < b.N; i++ {
		popcount.PopCountClearLowestBit(uint64(n))
	}
}
