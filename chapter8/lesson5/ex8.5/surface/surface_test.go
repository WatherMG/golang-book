package surface

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
	for _, workers := range []int{1, 2, 4, 6, 8, 10, 12, 16, runtime.GOMAXPROCS(-1), 24, 32, 64, 128, 256} {
		b.Run(fmt.Sprintf("workers=%d", workers), func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(width * height * 4)
			for i := 0; i < b.N; i++ {
				ConcurrentRender(workers)
			}
		})
	}
}
