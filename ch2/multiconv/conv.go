package multiconv

// CToF converts Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// FToM converts Feets to Meters
func FToM(ft Feet) Meter {
	return Meter(ft / 3.2808)
}

// MToF converts Meter to Feet
func MToF(m Meter) Feet {
	return Feet(m * 3.2808)
}

// KToP converts Kilo to Pound
func KToP(kg Kilo) Pound {
	return Pound(kg / 2.2046)
}

// PToK converts Pound to Kilo
func PToK(p Pound) Kilo {
	return Kilo(p * 2.2046)
}

// KToC converts Kelvin to Celsius
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}
