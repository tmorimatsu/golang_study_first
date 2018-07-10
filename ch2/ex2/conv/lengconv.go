package conv

import "fmt"

type Meter float64
type Feet float64

const errMsgLength string = "長さの取りうる範囲ではありません"

func (m Meter) String() string {
	validationLength(float64(m))
	return fmt.Sprintf("%gM", m)
}
func (f Feet) String() string {
	validationLength(float64(f))
	return fmt.Sprintf("%gF", f)
}

func MeterToFeet(m Meter) Feet {
	validationLength(float64(m))
	return Feet(m * 3.28084)
}
func FeetToMeter(f Feet) Meter {
	validationLength(float64(f))
	return Meter(f / 3.28084)
}

func validationLength(f float64) {
	if f < 0 {
		panic(errMsgLength)
	}
}
