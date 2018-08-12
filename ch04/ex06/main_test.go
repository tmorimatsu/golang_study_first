package main

import "testing"

func TestCompressSpaces(t *testing.T) {
	t.Log("隣接するスペースが1つに圧縮される")
	expected := "abcd ef g"
	actual := compressSpaces([]byte("abcd   ef     g"))
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
