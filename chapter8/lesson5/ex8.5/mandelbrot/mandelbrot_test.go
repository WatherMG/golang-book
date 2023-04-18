package mandelbrot

import (
	"fmt"
	"runtime"
	"testing"
)

func BenchmarkSerialRender(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(width * height * 4)
	for i := 0; i < b.N; i++ {
		SerialRender()
	}
}

func BenchmarkConcurrentRender(b *testing.B) {
	for _, workers := range []int{2, 4, 6, 8, 10, 12, 16, runtime.GOMAXPROCS(-1), 24, 32, 64, 128, 256} {
		b.Run(fmt.Sprintf("workers=%d", workers), func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(width * height * 4)
			for i := 0; i < b.N; i++ {
				ConcurrentRender(workers)
			}
		})
	}
}

/*
goos: windows
goarch: amd64
pkg: GolangBook/chapter8/lesson5/ex8.5/mandelbrot
cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
BenchmarkSerialRender-16                               8         135229912 ns/op          31.02 MB/s     8389131 B/op    1048579 allocs/op
BenchmarkConcurrentRender/workers=1-16                 8         134687112 ns/op          31.14 MB/s     8398728 B/op    1048583 allocs/op
BenchmarkConcurrentRender/workers=2-16                15          68737407 ns/op          61.02 MB/s     8398841 B/op    1048584 allocs/op
BenchmarkConcurrentRender/workers=4-16                33          36565912 ns/op         114.71 MB/s     8398915 B/op    1048586 allocs/op
BenchmarkConcurrentRender/workers=6-16                48          25412958 ns/op         165.05 MB/s     8399071 B/op    1048588 allocs/op
BenchmarkConcurrentRender/workers=8-16                58          21586160 ns/op         194.31 MB/s     8399244 B/op    1048591 allocs/op
BenchmarkConcurrentRender/workers=10-16               62          20416432 ns/op         205.44 MB/s     8399163 B/op    1048593 allocs/op
BenchmarkConcurrentRender/workers=12-16               64          18495567 ns/op         226.77 MB/s     8399466 B/op    1048596 allocs/op
BenchmarkConcurrentRender/workers=16-16               74          16619359 ns/op         252.37 MB/s     8399441 B/op    1048600 allocs/op
BenchmarkConcurrentRender/workers=16#01-16            72          16842612 ns/op         249.03 MB/s     8399800 B/op    1048600 allocs/op
BenchmarkConcurrentRender/workers=24-16               68          16830746 ns/op         249.20 MB/s     8399595 B/op    1048608 allocs/op
BenchmarkConcurrentRender/workers=32-16               74          16581818 ns/op         252.95 MB/s     8399811 B/op    1048615 allocs/op
BenchmarkConcurrentRender/workers=64-16               75          16557548 ns/op         253.32 MB/s     8400969 B/op    1048647 allocs/op
BenchmarkConcurrentRender/workers=128-16              72          16725050 ns/op         250.78 MB/s     8403126 B/op    1048711 allocs/op
BenchmarkConcurrentRender/workers=256-16              75          16694523 ns/op         251.24 MB/s     8407579 B/op    1048843 allocs/op
*/
