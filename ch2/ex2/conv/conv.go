package conv

import "math"

/*
func CelsiusToFahrenheit(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FahrenheitToCelsius(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
*/

func round(f float64) float64 {
	return math.Floor(f + .5)
}
