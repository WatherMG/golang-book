/*
Countdown реализует обратный отсчет для запуска ракеты.
*/

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// ...создаем канал abort...
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // Читаем 1 байт
		abort <- struct{}{}
	}()

	fmt.Println("Начинаю отсчет. Нажмите <ENTER> для отмены.")
	select {
	case <-time.After(10 * time.Second):
	// Ничего не делаем
	case <-abort:
		fmt.Println("Запуск отменен!")
		return
	}
	launch()
}

func launch() {
	fmt.Println("Запуск!")
}
