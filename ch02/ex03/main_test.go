package main

import (
	"testing"

	"./popcount"
)

const n uint64 = 114514

// default
func TestPopCount(t *testing.T) {
	expected := 11
	actual := popcount.PopCount(n)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(uint64(n))
	}
}

// ch2/ex3
func TestPopCountLoop(t *testing.T) {
	expected := 11
	actual := popcount.PopCountLoop(n)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(uint64(n))
	}
}

// ch2/ex4
func TestPopCountBitShift64(t *testing.T) {
	expected := 11
	actual := popcount.PopCountBitShift64(n)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
func BenchmarkPopCountBitShift64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountBitShift64(uint64(n))
	}
}

// ch2/ex5
func TestPopCountClearLowestBit(t *testing.T) {
	expected := 11
	actual := popcount.PopCountClearLowestBit(n)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
func BenchmarkPopCountClearLowestBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountClearLowestBit(uint64(n))
	}
}
