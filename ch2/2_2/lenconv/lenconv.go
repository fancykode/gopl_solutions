// Package lenconv meters and feet conversions.
package lenconv

import "fmt"

type Meter float64
type Foot float64

// FtToMt converts feet to meters.
func FtToMt(f Foot) Meter {
	return Meter(f * 0.3048)
}

// MtToFt converts meters to feet.
func MtToFt(m Meter) Foot {
	return Foot(m * 3.28084)
}

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (f Foot) String() string  { return fmt.Sprintf("%gft", f) }
