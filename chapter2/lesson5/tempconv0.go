/*
Example 2-5-1
Пакет tempconv выполняет вычисления температур по Цельсию (Celsius) и по Фаренгейту (Fahrenheit),
*/

package tempconv

import "fmt"

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
