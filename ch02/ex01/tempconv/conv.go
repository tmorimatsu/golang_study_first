package tempconv

func CToF(c Celsius) Fahrenheit {
	validationCelsius(c)
	return Fahrenheit(c*9/5 + 32)
}
func FToC(f Fahrenheit) Celsius {
	validationFahrenheit(f)
	return Celsius((f - 32) * 5 / 9)
}
func FToK(f Fahrenheit) Kelvin {
	validationFahrenheit(f)
	return CToK(FToC(f))
}
func CToK(c Celsius) Kelvin {
	validationCelsius(c)
	return Kelvin(c + 273.15)
}
func KToC(k Kelvin) Celsius {
	validationKelvin(k)
	return Celsius(k - 273.15)
}
func KToF(k Kelvin) Fahrenheit {
	validationKelvin(k)
	return CToF(KToC(k))
}

func validationCelsius(c Celsius) {
	if c+273.15 < 0 {
		panic(errmsg)
	}
}

func validationFahrenheit(f Fahrenheit) {
	if ((f-32)*5/9)+273.15 < 0 {
		panic(errmsg)
	}
}

func validationKelvin(k Kelvin) {
	if k < 0 {
		panic(errmsg)
	}
}
