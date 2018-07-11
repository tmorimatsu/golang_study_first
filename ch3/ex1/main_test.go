package main

import (
	"testing"
)

func TestValidateInfinityFloat(t *testing.T) {
	expected := true
	actual := ValidateInfinityFloat()
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
