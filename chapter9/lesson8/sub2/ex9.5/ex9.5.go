package main

import (
	"fmt"
	"time"
)

func main() {
	pings := make(chan string)
	pongs := make(chan string)
	var i int

	start := time.Now()
	go func() {
		for {
			i++
			pings <- "ping"
			<-pongs
		}
	}()

	go func() {
		for {
			i++
			<-pings
			pongs <- "pong"
		}
	}()

	<-time.After(10 * time.Second)
	elapsed := time.Since(start)
	fmt.Printf("%.f op/s. i=%d, t=%s", float64(i)/elapsed.Seconds(), i, elapsed)

}

// 5597927 op/s. i=55984622, t=10.0009426s - string ["ping", "pong"]
// 5844477 op/s. i=58445279, t=10.0000721s - string ["ping", "pong"], buffer

// 5755891 op/s. i=57560415, t=10.0002606s - int 1
// 5810025 op/s. i=58103794, t=10.0005883s - int 1, buffer

// 5770534 op/s. i=57709983, t=10.0007928s - int i
// 5801099 op/s. i=58012627, t=10.0002682s - int i, buffer
