package bank

import (
	"sync"
	"testing"
)

func TestWithdrawConcurrent(t *testing.T) {
	Deposit(10000)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(amount int) {
			wg.Done()
			Withdraw(amount)
		}(i)
	}
	wg.Wait()

	if got, want := Balance(), 5050; got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestWithdraw(t *testing.T) {
	Deposit(50)
	b1 := Balance()
	ok := Withdraw(50)
	if !ok {
		t.Errorf("ok = false, want true. balance = %d", Balance())
	}
	expected := b1 - 50
	if b2 := Balance(); b2 != expected {
		t.Errorf("balance = %d, want %d", b2, expected)
	}
}

func TestWithdrawFailsIfInsufficientFunds(t *testing.T) {
	b1 := Balance()
	ok := Withdraw(b1 + 1)
	b2 := Balance()
	if ok {
		t.Errorf("ok = true, want false. balance = %d", b2)
	}
	if b2 != b1 {
		t.Errorf("balance = %d, want %d", b2, b1)
	}
}
