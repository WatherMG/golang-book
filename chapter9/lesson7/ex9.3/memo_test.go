package memo

import (
	"testing"
	"time"

	"GolangBook/chapter9/lesson7/ex9.3/memotest"
)

const timeout = 400 * time.Millisecond

var httpGetBody = memotest.HTTPGetBody

func TestSequential(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	// testing cancellation
	done := make(chan struct{})
	go func() {
		time.Sleep(timeout)
		close(done)
	}()
	memotest.Sequential(t, m, done)
}

func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	done := make(chan struct{})
	memotest.Concurrent(t, m, done)
}

func TestCancel(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	done := make(chan struct{})
	memotest.ConcurrentCancel(t, m, done)
}

/*
go test -race -run=TestConcurrent
https://godoc.org, 516.204207ms, 295 bytes
https://golang.org, 880.118212ms, 1579 bytes
http://gopl.io, 2.771554139s, 4154 bytes
https://godoc.org, 2.771545839s, 295 bytes
https://golang.org, 2.771534339s, 1579 bytes
https://play.golang.org, 3.193311445s, 1579 bytes
http://gopl.io, 3.193274845s, 4154 bytes
https://play.golang.org, 3.193272845s, 1579 bytes
PASS
*/
