package main

import (
	"crypto/sha256"
	"testing"
)

func TestHammingdistanceOK(t *testing.T) {
	v := sha256.Sum256([]byte("X"))
	expected := 0
	actual := hammingdistance(v, v)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestHammingdistanceDiff(t *testing.T) {
	v1 := sha256.Sum256([]byte("x"))
	v2 := sha256.Sum256([]byte("X"))
	expected := 125
	actual := hammingdistance(v1, v2)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
