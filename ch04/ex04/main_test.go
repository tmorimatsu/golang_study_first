package main

import "testing"

func TestRotate(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rotatedArr := rotate(arr, 3)
	expected := arr[3]
	actual := rotatedArr[0]
	if expected != actual {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
