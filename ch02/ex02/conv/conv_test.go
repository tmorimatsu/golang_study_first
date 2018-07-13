package conv

import "testing"

func TestMeterToFeet(t *testing.T) {
	expected := Feet(32.81)
	actual, err := MeterToFeet(10)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestMeterToFeetError(t *testing.T) {
	expected, err := MeterToFeet(-1)
	if !(expected == -1 && err != nil){
		t.Errorf("エラーです"/*"actual %v\nwant %v", actual, expected*/)
	}
	/*actual := MeterToFeet(10)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}*/
}

func TestFeetToMeter(t *testing.T) {
	expected := Meter(3.05)
	actual := FeetToMeter(10)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
