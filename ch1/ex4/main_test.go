package main

import "testing"

// main関数から書き直さないと動かなさそう

func TestCountLines(t *testing.T) {
	expected := 21
	actual := CountLines("test_text.txt")
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}