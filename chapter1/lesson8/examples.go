package main

import "fmt"

// Инструкция switch
func s() {
	x := 3
	switch x {
	case 1:
		fmt.Println("x is equal to 1")
	case 2:
		fmt.Println("x is equal to 2")
	case 3:
		fmt.Println("x is equal to 3")
	default:
		fmt.Println("x is not equal to 1, 2, or 3")
	}
	// Вывод: "x is equal to 3"
}

// Инструкции break и continue:
func b() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue // Пропускает итерацию цикла при i == 3
		}
		if i == 8 {
			break // Завершает цикл при i == 8
		}
		fmt.Println(i)
	}
	// Вывод: "0\n1\n2\n4\n5\n6\n7"
}

// Объявление type:
func t() {
	type Celsius float64 // Объявление нового типа Celsius

	var temperature Celsius = 20.0 // Использование нового типа

	fmt.Printf("The temperature is %.1f degrees Celsius.\n", temperature)
	// Вывод: "The temperature is 20.0 degrees Celsius."
}

// Переключатель без тегов:
func sn() {
	x := 3
	switch {
	case x < 0:
		fmt.Println("x is negative")
	case x == 0:
		fmt.Println("x is zero")
	case x > 0:
		fmt.Println("x is positive")
	}
	// Вывод: "x is positive"
}

// Инструкция goto:
func g() {
	i := 0
loop: // Метка для инструкции goto
	if i < 10 {
		fmt.Println(i)
		i++
		goto loop // Переход к метке loop
	}
	// Вывод: "0\n1\n2\n3\n4\n5\n6\n7\n8\n9"
}
