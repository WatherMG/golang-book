/*
Exercise 5.15
Напишите вариативные функции max и min, аналогичные функции sum.
Что должны делать эти функции, будучи вызванными без аргументов?
Напишите варианты функций, требующие как минимум одного аргумента.
*/

package main

import "fmt"

func main() {
	fmt.Println(max(-444, -1, -2, -3, -4, -5, -111, -333))
	fmt.Println(min(-444, -1, -2, -3, -4, -5, -111, -333))
	fmt.Println(max2(1, -1, -2, -3, -4, -5, -111, -333))
	fmt.Println(min2(10, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(max2(1))
	fmt.Println(min2(10))
}

func max(values ...int) (m int) {
	if len(values) == 0 {
		panic("max: must have at least one argument")
	}
	m = values[0]
	for _, num := range values {
		if num > m {
			m = num
		}
	}
	return m
}

func min(values ...int) (m int) {
	if len(values) == 0 {
		panic("min: must have at least one argument")
	}
	m = values[0]
	for _, num := range values {
		if num < m {
			m = num
		}
	}
	return m
}

func max2(first int, values ...int) (m int) {
	m = first
	for _, num := range values {
		if num > m {
			m = num
		}
	}
	return m
}

func min2(first int, values ...int) (m int) {
	m = first
	for _, num := range values {
		if num < m {
			m = num
		}
	}
	return m
}
