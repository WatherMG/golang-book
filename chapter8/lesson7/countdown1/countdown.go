/*
Countdown реализует обратный отсчет для запуска ракеты.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Начинаю отсчет.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}

func launch() {
	fmt.Println("Запуск!")
}
