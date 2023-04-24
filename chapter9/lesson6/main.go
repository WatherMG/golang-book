package main

import (
	"fmt"
	"sync"
)

var counter int
var wg sync.WaitGroup

// var mu sync.Mutex // блокирует горутины, пока критический участок кода занят другой горутиной

func increment() {
	// mu.Lock()
	// defer mu.Unlock()
	counter++
	wg.Done()
}

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment()
	}
	wg.Wait()
	fmt.Printf("Counter value: %d\n", counter)
}

/*
go run -race main.go

WARNING: DATA RACE
Write at 0x00c00008a010 by goroutine 10:
  main.increment()
      /home/user/main.go:12 +0x3a

Previous write at 0x00c00008a010 by goroutine 9:
  main.increment()
      /home/user/main.go:12 +0x3a

Goroutine 10 (running) created at:
  main.main()
      /home/user/main.go:18 +0x73

Goroutine 9 (finished) created at:
  main.main()
      /home/user/main.go:18 +0x73
==================
Counter value: 97
Found 1 data race(s)
exit status 66
*/
