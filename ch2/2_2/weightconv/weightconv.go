// Package weightconv performs kilograms and pounds conversions.
package weightconv

import "fmt"

type Kilogram float64
type Pound float64

// PndToKg converts pounds to kilograms.
func PndToKg(p Pound) Kilogram {
	return Kilogram(p * 0.45359237)
}

// KgToPnd converts kilograms to pounds.
func KgToPnd(k Kilogram) Pound {
	return Pound(k * 2.205)
}

func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
func (p Pound) String() string    { return fmt.Sprintf("%glb", p) }
