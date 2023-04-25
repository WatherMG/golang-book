package pipeline

import (
	"testing"
)

func TestPipeline(t *testing.T) {
	in, out := pipeline(3)
	in <- 1
	// Это значение передается по 3 каналам
	// (исключая главную горутину и каналы in out) с помощью 3 горутин
	t.Log(<-out)
}

func bench(b *testing.B, stages int) {
	b.Helper()
	in, out := pipeline(stages)
	for i := 0; i < b.N; i++ {
		go func() {
			in <- 1
		}()
		<-out
	}
	close(in)
}

func BenchmarkPipeline1(b *testing.B) {
	bench(b, 1)
}

func BenchmarkPipeline512(b *testing.B) {
	bench(b, 512)
}

func BenchmarkPipeline1024(b *testing.B) {
	bench(b, 1024)
}

func BenchmarkPipeline2048(b *testing.B) {
	bench(b, 2048)
}

func BenchmarkPipeline25k(b *testing.B) {
	bench(b, 25000)
}

func BenchmarkPipeline50k(b *testing.B) {
	bench(b, 50000)
}

func BenchmarkPipeline100k(b *testing.B) {
	bench(b, 100000)
}
func BenchmarkPipeline500k(b *testing.B) {
	bench(b, 500000)
}
func BenchmarkPipeline760k(b *testing.B) {
	bench(b, 760000)
}
func BenchmarkPipeline100kk(b *testing.B) {
	bench(b, 1000000)
}
func BenchmarkPipeline200kk(b *testing.B) {
	bench(b, 2000000)
}
func BenchmarkPipeline500kk(b *testing.B) {
	bench(b, 5000000)
}
func BenchmarkPipeline1000k(b *testing.B) {
	bench(b, 10000000) // error
}
