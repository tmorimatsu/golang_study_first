package conv

import (
	_ "errors"
	"fmt"
)

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

func MeterToFeet(m Meter) /*(Feet, error)*/ Feet {
	if !validationLength(float64(m)) {
		//return -1, errors.New(errMsgLength)
		panic(errMsgLength)
	}
	return Feet(Round(float64(m*3.28084), 2)) //, nil
}
func FeetToMeter(f Feet) /*(Meter, error)*/ Meter {
	if !validationLength(float64(f)) {
		//return -1, errors.New(errMsgLength)
		panic(errMsgLength)
	}
	return Meter(Round(float64(f/3.28084), 2)) //, nil
}

func validationLength(f float64) bool {
	if f < 0 {
		//panic(errMsgLength)
		return false
	}
	return true
}

/*
panicではなくこっちを使う？
	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
	os.Exit(1)
*/
