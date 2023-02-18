package mandelbrot

import (
	"image/color"
	"testing"
)

func bench(b *testing.B, f func(complex128) color.Color) {
	for i := 0; i < b.N; i++ {
		f(complex(float64(i), float64(i)))
	}
}

func BenchmarkMandelbrotComplex64(b *testing.B) {
	bench(b, mandelbrot64)
}

func BenchmarkMandelbrotComplex128(b *testing.B) {
	bench(b, mandelbrot128)
}

func BenchmarkMandelbrotBigFloat(b *testing.B) {
	bench(b, mandelbrotBigFloat)
}

func BenchmarkMandelbrotBigRat(b *testing.B) {
	bench(b, mandelbrotBigRat)
}
