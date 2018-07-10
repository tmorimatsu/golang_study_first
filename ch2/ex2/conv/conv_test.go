package conv

import "testing"

func TestMeterToFeet(t *testing.T) {
	expected := Feet(131.2336)
	actual := MeterToFeet(40)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
