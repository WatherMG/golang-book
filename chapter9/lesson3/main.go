package bank

import (
	"sync"
)

var mu sync.RWMutex
var balance int

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance += amount
}

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}
