/*
Мой пример для lesson13
*/

package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

const DEFAULT = "no case for type"

// Определение типа значения интерфейса.
func printType(x interface{}) string {
	// Используем type switch для определения типа значения x.
	switch x.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case float64:
		return "float64"
	case float32:
		return "float32"
	default:
		return DEFAULT
	}
}

// Обработка значений разных типов.
func printValue(x interface{}) string {
	// Используем расширенную форму type switch с переменной x.
	// В каждом case значение x извлекается и присваивается новой переменной x.
	// В case int переменная x имеет тип int, а в case string - тип string.
	switch x := x.(type) {
	case int:
		return fmt.Sprintf("%d", x)
	case uint:
		return fmt.Sprintf("%d", x)
	case string:
		return x
	case bool:
		return fmt.Sprintf("%t", x)
	case float64:
		return fmt.Sprintf("%.2f", x)
	case float32:
		return fmt.Sprintf("%.2f", x)
	default:
		return DEFAULT
	}
}

// Объединение нескольких case.
func printTypeOr(x interface{}) string {
	// Используем запятую для объединения нескольких case.
	switch x := x.(type) {
	case int, uint:
		return fmt.Sprintf("%T: %[1]d", x)
	case string:
		return fmt.Sprintf("%T: %[1]s", x)
	case bool:
		return fmt.Sprintf("%T: %[1]t", x)
	case float32, float64:
		return fmt.Sprintf("%T: %.2[1]f", x)
	case []byte:
		return fmt.Sprintf("%T: %[1]b=%[1]d", x)
	default:
		return DEFAULT
	}
}

func main() {
	const format = "%v\t%v\t%v\t%v\n"

	values := []interface{}{42, "hello", true, 3.14, float32(4.13), uint(24), []byte{10, 20}, complex(1.0, -2)}

	tw := new(tabwriter.Writer).Init(os.Stdout, 2, 16, 2, ' ', 0)
	fmt.Fprintf(tw, format, "#", "Type", "Value", "Type and value")
	fmt.Fprintf(tw, format, "-", "----", "-----", "--------------")

	for i, v := range values {
		fmt.Fprintf(tw, format, i+1, printType(v), printValue(v), printTypeOr(v))
	}
	tw.Flush()
}
