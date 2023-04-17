/*
Example 8.9
Pipeline1 демонстрирует бесконечный трехступенчатый конвейер.
*/
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Генерация
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Возведение в квадрат
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break // Канал закрыт и опустошен
			}
			squares <- x * x
		}
		close(squares)
	}()

	// Вывод (в главной горутине)
	for {
		fmt.Println(<-squares)
	}
}
