package equal

import "testing"

func TestAlmostEqual(t *testing.T) {
	ts := []struct {
		v1       float64
		v2       float64
		expected bool
	}{
		{100, 100, true},
		{100.00000002, 100, false},
		{100.0000001, 100, false},
		{-100.0000000001, -100, true},
		{100.000000001, 100, true},
	}

	for _, k := range ts {
		actual := Equal(k.v1, k.v2)
		if k.expected != actual {
			t.Fatalf("exptected %v, but actual %v", k.expected, actual)
		}
	}
}
