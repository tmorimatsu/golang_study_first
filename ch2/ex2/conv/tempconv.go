package conv

import "fmt"

type Celsius float64
type Fahrenheit float64

// 定義の追加
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// 単位の追加
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

// 単位の変換
func CelsiusToFahrenheit(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FahrenheitToCelsius(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
