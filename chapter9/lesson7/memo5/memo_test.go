package memo_test

import (
	"testing"

	"GolangBook/chapter9/lesson7/memo5"
	"GolangBook/chapter9/lesson7/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func TestSequential(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
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
