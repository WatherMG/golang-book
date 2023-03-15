/*
Example 5.14
*/

package main

import (
	"log"
	"time"
)

func main() {

	BigSlowOperation()
}

func BigSlowOperation() {
	defer trace("bigSlowOperation")() // Не забываем о скобках!
	// ...длительная работа...
	time.Sleep(10 * time.Second) // Имитация долгой работы
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("вход в %s", msg)
	return func() {
		log.Printf("выход из %s (%s)", msg, time.Since(start))
	}
}
