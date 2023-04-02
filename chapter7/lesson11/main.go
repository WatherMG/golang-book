package main

import (
	"fmt"
	"os"
)

func main() {
	// Пытаемся открыть файл, который не существует
	_, err := os.Open("/no/such/file")
	if err != nil {
		// С помощью декларации типа проверяем тип ошибки
		if pe, ok := err.(*os.PathError); ok {
			// Если тип ошибки == PathError выводим подробное сообщение об ошибке
			fmt.Printf("Операция: %s\nФайл: %s\nОшибка: %v\n", pe.Op, pe.Path, pe.Err)
		} else {
			// Иначе это сообщение
			fmt.Println("Это другая ошибка:", err)
		}
	}
}
