/*
Exercise 2-6-1 (2.2)
The lenghtconv package performs length conversions - feet and meters
*/

package lenghtconv

import "fmt"

type Meter float64
type Feet float64

const (
	LengthCoefficient = 3.28084
)

func (m Meter) String() string {
	return fmt.Sprintf("%.2f m", m)
}

func (f Feet) String() string {
	return fmt.Sprintf("%.2f ft", f)
}
