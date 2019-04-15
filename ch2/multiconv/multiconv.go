// Package multiconv makes conversion between Celsius and Fahrenheit, Meter and Feet, Kilo and Pound
package multiconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Feet float64
type Meter float64
type Pound float64
type Kilo float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}

func (ft Feet) String() string {
	return fmt.Sprintf("%gf", ft)
}

func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}

func (p Pound) String() string {
	return fmt.Sprintf("%gp", p)
}

func (kg Kilo) String() string {
	return fmt.Sprintf("%gkg", kg)
}
