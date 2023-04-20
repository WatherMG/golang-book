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
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Начинаю отсчет. Нажмите <ENTER> для отмены.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// ничего не делаем
		case <-abort:
			fmt.Println("Запуск отменен!")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Запуск!")
}
