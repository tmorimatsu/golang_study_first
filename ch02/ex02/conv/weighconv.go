package conv

import "fmt"

type Kilogram float64
type Pond float64

const errMsgWeigh string = "重さの取りうる範囲ではありません"

func (k Kilogram) String() string {
	validationWeigh(float64(k))
	return fmt.Sprintf("%gkg", k)
}
func (p Pond) String() string {
	validationWeigh(float64(p))
	return fmt.Sprintf("%glb", p)
}

func KilogramToPond(k Kilogram) Pond {
	validationWeigh(float64(k))
	return Pond(k / 0.453592)
}
func PondToKilogram(p Pond) Kilogram {
	validationWeigh(float64(p))
	return Kilogram(p * 0.453592)
}

func validationWeigh(f float64) {
	if f < 0 {
		panic(errMsgWeigh)
	}
}
