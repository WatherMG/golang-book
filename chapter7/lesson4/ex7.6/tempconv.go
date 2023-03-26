/*
Exercise 7.6
Добавьте в tempflag поддержку температуры по шкале Кельвина.
*/

package tempconv

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout

type Celsius float64
type Fahrenheit float64
type Kelvin float64

// FToC преобразует температуру по Фаренгейту в температуру по Цельсию.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

// KToC преобразует температуру по Кельвину в температуру по Цельсию
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// Реализуем метод String для всех типов
func (c Celsius) String() string { return fmt.Sprintf("%.2f°C", c) }

// *celsiusFlag соответствует интерфейсу flag.Value
type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("неверная температура %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
