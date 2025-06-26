// Package tempconv performs Celsius, Kelvin and Fahrenheit conversions.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	BoilingK      Kelvin  = 373.15
	FreezingK     Kelvin  = 273.15
	AbsoluteZeroK Kelvin  = 0
)

func (c Celsius) String() string    { return fmt.Sprintf("%gC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gF", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrehnheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts a Celsius temperature to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// KToC converts a Kelvin temperature to Celsius
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }
