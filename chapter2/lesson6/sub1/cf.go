/*
Example 2-6-3
Cf конвертирует числовой аргумент в температуру по Цельсию и по Фаренгейту,
*/
package main

import (
	"fmt"
	"os"
	"strconv"

	"GolangBook/chapter2/lesson6/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] { // Получаем список аргументов
		t, err := strconv.ParseFloat(arg, 64) // приводим к типу float64 аргумент, который является string
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t) // приводим к типу Fahrenheit переменную t (имеет базовый тип float64)
		c := tempconv.Celsius(t)    // приводим к типу Celsius переменную t (имеет базовый тип float64)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c)) // выводим текст, используя преобразование температур
	}
}
