package mandelbrot

import (
	"runtime"
	"testing"
)

func BenchmarkSerialRender(b *testing.B) {
	// b.ReportAllocs()
	// b.SetBytes(width * height * 4)
	for i := 0; i < b.N; i++ {
		SerialRender()
	}
}

/*func BenchmarkConcurrentRender(b *testing.B) {
	for _, workers := range []int{2, 4, 6, 8, 10, 12, 16, runtime.GOMAXPROCS(-1), 24, 32, 64, 128, 256} {
		b.ReportAllocs()
		b.Run(fmt.Sprintf("workers=%d", workers), func(b *testing.B) {
			b.SetBytes(width * height * 4)
			for i := 0; i < b.N; i++ {
				ConcurrentRender(workers)
			}
		})
	}
}*/

func benchConcurrentRender(b *testing.B, procs int) {
	b.Helper()
	for i := 0; i < b.N; i++ {
		ConcurrentRender(procs)
	}
}

func Benchmark1(b *testing.B) {
	benchConcurrentRender(b, 1)
}

func BenchmarkMax(b *testing.B) {
	benchConcurrentRender(b, runtime.GOMAXPROCS(-1))
}
func Benchmark8(b *testing.B) {
	benchConcurrentRender(b, 8)
}
func Benchmark16(b *testing.B) {
	benchConcurrentRender(b, 16)
}

func Benchmark32(b *testing.B) {
	benchConcurrentRender(b, 32)
}

func Benchmark64(b *testing.B) {
	benchConcurrentRender(b, 64)
}

func Benchmark128(b *testing.B) {
	benchConcurrentRender(b, 128)
}

/*
GOMAXPROCS=1 go test -bench=.
BenchmarkSerialRender          8         134254076 ns/op
Benchmark1                     8         135635901 ns/op
BenchmarkMax                   8         134874538 ns/op
Benchmark8                     8         134862776 ns/op
Benchmark16                    8         135587513 ns/op
Benchmark32                    8         135236638 ns/op
Benchmark64                    8         135149263 ns/op
Benchmark128                   8         135161538 ns/op

GOMAXPROCS=2 go test -bench=.
BenchmarkSerialRender-2                8         134239838 ns/op
Benchmark1-2                           8         133562513 ns/op
BenchmarkMax-2                        16          69794613 ns/op
Benchmark8-2                          16          69896638 ns/op
Benchmark16-2                         16          69305613 ns/op
Benchmark32-2                         16          69507013 ns/op
Benchmark64-2                         16          68914225 ns/op
Benchmark128-2                        16          71332375 ns/op

GOMAXPROCS=4 go test -bench=.
BenchmarkSerialRender-4                8         134157626 ns/op
Benchmark1-4                           8         134278950 ns/op
BenchmarkMax-4                        28          37029761 ns/op
Benchmark8-4                          33          37313843 ns/op
Benchmark16-4                         32          37201316 ns/op
Benchmark32-4                         32          37597434 ns/op
Benchmark64-4                         32          38152953 ns/op
Benchmark128-4                        33          38172836 ns/op

GOMAXPROCS=8 go test -bench=.
BenchmarkSerialRender-8                8         133731876 ns/op
Benchmark1-8                           8         134376638 ns/op
BenchmarkMax-8                        63          20437670 ns/op
Benchmark8-8                          58          20750116 ns/op
Benchmark16-8                         55          20753089 ns/op
Benchmark32-8                         55          21229847 ns/op
Benchmark64-8                         52          21910100 ns/op
Benchmark128-8                        60          22703653 ns/op

GOMAXPROCS=16 go test -bench=.
BenchmarkSerialRender-16               8         134044526 ns/op
Benchmark1-16                          8         134213952 ns/op
BenchmarkMax-16                       75          16864347 ns/op
Benchmark8-16                         66          21419043 ns/op
Benchmark16-16                        67          16827097 ns/op
Benchmark32-16                        69          17229092 ns/op
Benchmark64-16                        70          17958820 ns/op
Benchmark128-16                       61          18787707 ns/op

GOMAXPROCS=32 go test -bench=.
BenchmarkSerialRender-32               8         134886552 ns/op
Benchmark1-32                          8         136213315 ns/op
BenchmarkMax-32                       66          17402223 ns/op
Benchmark8-32                         56          21616943 ns/op
Benchmark16-32                        72          16817653 ns/op
Benchmark32-32                        74          16913852 ns/op
Benchmark64-32                        69          17624924 ns/op
Benchmark128-32                       67          18565581 ns/op

GOMAXPROCS=64 go test -bench=.
BenchmarkSerialRender-64               8         134269876 ns/op
Benchmark1-64                          8         134447788 ns/op
BenchmarkMax-64                       72          18120825 ns/op
Benchmark8-64                         66          21799911 ns/op
Benchmark16-64                        66          17336691 ns/op
Benchmark32-64                        74          17390031 ns/op
Benchmark64-64                        64          17955497 ns/op
Benchmark128-64                       61          19320900 ns/op
*/
