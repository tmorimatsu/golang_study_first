package main

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		values []int
		want   int
	}{
		{[]int{1, 2, 0, 4, -5}, 4},
		{[]int{-5, -4, -3, -2, -1}, -1},
		{[]int{}, 0},
	}

	for _, test := range tests {
		actual, err := max(test.values...)
		if actual != test.want {
			t.Errorf("expected (%d, nil) but got (%d, %v)", err, test.want, actual)
		}
		if len(test.values) == 0 {
			if err == nil {
				t.Errorf("expected (%d, %v) but got (%d, nil)", test.want, err, actual)
			}
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		values []int
		want   int
	}{
		{[]int{1, 2, 0, 4, 5}, 0},
		{[]int{-5, -4, -3, -2, -1}, -5},
		{[]int{}, 0},
	}

	for _, test := range tests {
		actual, err := min(test.values...)
		if actual != test.want {
			t.Errorf("expected (%d, nil) but got (%d, %v)", err, test.want, actual)
		}
		if len(test.values) == 0 {
			if err == nil {
				t.Errorf("expected (%d, %v) but got (%d, nil)", test.want, err, actual)
			}
		}
	}
}
