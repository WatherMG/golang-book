/*
Пакет tempconv выполняет преобразования температур
*/

package tempconv

import "fmt"

/*
Exercise 2-5-1
Добавьте в пакет tempconv типы, константы и функции для работы с температурой по шкале Кельвина, в которой нуль градусов
соответствует температуре-273.15°С, а разница температур в 1К имеет ту же величину, что и 1°С
*/

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k-Kelvin(AbsoluteZeroC)) }
