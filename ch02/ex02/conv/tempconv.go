package conv

import "fmt"

type Celsius float64
type Fahrenheit float64

// 定義の追加
const (
	AbsoluteZeroC     Celsius = -273.15
	FreezingC         Celsius = 0
	BoilingC          Celsius = 100
	errMsgTemparature string  = "温度の取りうる範囲ではありません"
)

// 単位の追加
func (c Celsius) String() string {
	validationCelsius(c)
	return fmt.Sprintf("%g°C", c)
}
func (f Fahrenheit) String() string {
	validationFahrenheit(f)
	return fmt.Sprintf("%g°F", f)
}

// 単位の変換
func CelsiusToFahrenheit(c Celsius) Fahrenheit {
	validationCelsius(c)
	return Fahrenheit(Round(float64(c*9/5 + 32), 2))
}
func FahrenheitToCelsius(f Fahrenheit) Celsius {
	validationFahrenheit(f)
  return Celsius(Round(float64((f - 32) * 5 / 9), 2))
}

func validationCelsius(c Celsius) {
	if c+273.15 < 0 {
		panic(errMsgTemparature)
	}
}

func validationFahrenheit(f Fahrenheit) {
	if ((f-32)*5/9)+273.15 < 0 {
		panic(errMsgTemparature)
	}
}
