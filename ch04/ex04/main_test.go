package main

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		arr  []int
		i    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 3, []int{4, 5, 6, 7, 8, 1, 2, 3}},
	}
	for _, test := range tests {
		actual := rotate(test.arr, test.i)
		if !reflect.DeepEqual(actual, test.want) {
			t.Errorf("actual %v\nwant %v", actual, test.want)
		}
	}
}
