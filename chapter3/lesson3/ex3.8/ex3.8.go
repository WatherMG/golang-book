/*
Exercise 3.8
Mandelbrot создает PNG-изображение фрактала Мандельброта.
Визуализация фракталов при высоком разрешении требует высокой арифметической точности.
Реализуйте один и тот же фрактал с помощью четырех различных представлений чисел:
complex64, complexl28, big.Float и big.Rat. Сравните производительность и потребление памяти
при использовании разных типов. При каком уровне масштабирования артефакты визуализации становятся видимыми?
https://github.com/torbiak/gopl/blob/master/ex3.8/main.go
*/

package mandelbrot

import (
	"image/color"
	"math"
	"math/big"
	"math/cmplx"
)

func mandelbrot64(z complex128) color.Color {
	const iterations = 200

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
			switch {
			case n > 50:
				return color.RGBA{R: 100, A: 255}
			default:
				logScale := math.Log(float64(n)) / math.Log(float64(iterations))
				return color.RGBA{B: 255 - uint8(logScale*255), A: 255}
			}
		}
	}
	return color.Black
}

func mandelbrot128(z complex128) color.Color {
	const iterations = 200
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			switch {
			case n > 50:
				return color.RGBA{R: 100, A: 255}
			default:
				logScale := math.Log(float64(n)) / math.Log(float64(iterations))
				return color.RGBA{B: 255 - uint8(logScale*255), A: 255}
			}
		}
	}
	return color.Black
}

func mandelbrotBigFloat(z complex128) color.Color {
	const iterations = 200
	zR := (&big.Float{}).SetFloat64(real(z))
	zI := (&big.Float{}).SetFloat64(imag(z))
	var vR, vI = &big.Float{}, &big.Float{}
	for i := uint8(0); i < iterations; i++ {
		// v*v + z = (r+i)^2 = r^2 + 2ri + i^2
		vR2, vI2 := &big.Float{}, &big.Float{}
		vR2.Mul(vR, vR).Sub(vR2, (&big.Float{}).Mul(vI, vI)).Add(vR2, zR)
		vI2.Mul(vR, vI).Mul(vI2, big.NewFloat(2)).Add(vI2, zI)
		vR, vI = vR2, vI2
		squareSum := &big.Float{}
		squareSum.Mul(vR, vR).Add(squareSum, (&big.Float{}).Mul(vI, vI))
		if squareSum.Cmp(big.NewFloat(4)) == 1 {
			switch {
			case i > 50:
				return color.RGBA{R: 100, A: 255}
			default:
				logScale := math.Log(float64(i)) / math.Log(float64(iterations))
				return color.RGBA{B: 255 - uint8(logScale*255), A: 255}
			}
		}
	}
	return color.Black
}

func mandelbrotBigRat(z complex128) color.Color {
	// High-resolution images take an extremely long time to render with
	// iterations = 200. Multiplying arbitrary precision numbers has
	// algorithmic complexity of at least O(n*log(n)*log(log(n)))
	const iterations = 200
	zR := (&big.Rat{}).SetFloat64(real(z))
	zI := (&big.Rat{}).SetFloat64(imag(z))
	var vR, vI = &big.Rat{}, &big.Rat{}
	for i := uint8(0); i < iterations; i++ {
		// v*v + z = (r+i)^2 = r^2 + 2ri + i^2
		vR2, vI2 := &big.Rat{}, &big.Rat{}
		vR2.Mul(vR, vR).Sub(vR2, (&big.Rat{}).Mul(vI, vI)).Add(vR2, zR)
		vI2.Mul(vR, vI).Mul(vI2, big.NewRat(2, 1)).Add(vI2, zI)
		vR, vI = vR2, vI2
		squareSum := &big.Rat{}
		squareSum.Mul(vR, vR).Add(squareSum, (&big.Rat{}).Mul(vI, vI))
		if squareSum.Cmp(big.NewRat(4, 1)) == 1 {
			switch {
			case i > 50:
				return color.RGBA{R: 100, A: 255}
			default:
				logScale := math.Log(float64(i)) / math.Log(float64(iterations))
				return color.RGBA{B: 255 - uint8(logScale*255), A: 255}
			}
		}
	}
	return color.Black
}
