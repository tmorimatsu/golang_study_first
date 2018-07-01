package conv

import "fmt"

type Meter float64
type Feet float64

func (m Meter) String() string { return fmt.Sprintf("%gM", m) }
func (f Feet) String() string  { return fmt.Sprintf("%gF", f) }

func MeterToFeet(m Meter) Feet { return Feet(m * 3.28084) }
func FeetToMeter(f Feet) Meter { return Meter(f / 3.28084) }
