package main

import "testing"

func TestReverse(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	expected := arr[2]
	reverse(&arr)
	actual := arr[4]
	if expected != actual {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
