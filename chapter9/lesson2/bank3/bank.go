/*
Example 9.3
Package bank предоставляет безопасный с точки зрения
параллельности банк с одним счетом
*/

package main

import "sync"

var (
	mu      sync.Mutex
	banalce int
)

func Deposit(amount int) {
	mu.Lock()
	banalce = banalce + amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := banalce
	mu.Unlock()
	return b
}
