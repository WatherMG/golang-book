/*
Example 9.2
Package bank предоставляет безопасный с точки зрения
параллельности банк с одним счетом
*/

package main

var (
	sema    = make(chan struct{}, 1) // Бинарный семафор для защиты balance
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // Захват токена
	balance = balance + amount
	<-sema // Освобождение токена
}

func Balance() int {
	sema <- struct{}{} // Захват токена
	b := balance
	<-sema // Освобождение токена
	return b
}
