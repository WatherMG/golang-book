/*
Exercise 2-6-1 (2.2)
The weightconv package performs weight conversions of pounds and kilograms
*/

package weightconv

import "fmt"

type Kilos float64
type Pounds float64

const (
	KiloPoundCoefficient = 0.45359237
)

func (k Kilos) String() string  { return fmt.Sprintf("%.2f kg.", k) }
func (p Pounds) String() string { return fmt.Sprintf("%.2f lbs", p) }
