package conv

import "math"

func Round(f float64, places int) (float64) {
	shift := math.Pow(10, float64(places))
	return math.Floor(f * shift + .5) / shift
}
