/*
Exercise 3.7
Mandelbrot создает PNG-изображение фрактала Мандельброта.
Еще один простой фрактал использует метод Ньютона для поиска комплексных решений уравнения z*-\ = 0.
Закрасьте каждую точку цветом, соответствующим тому корню из четырех, которого она достигает, а интенсивность цвета
должна соответствовать количеству итераций, необходимых для приближения к этому корню.
https://github.com/torbiak/gopl/blob/master/ex3.7/main.go
*/

package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"os"
)

type Func func(complex128) complex128

var palettes = []color.RGBA{
	{170, 57, 57, 255},
	{170, 108, 57, 255},
	{34, 102, 102, 255},
	{45, 136, 45, 255},
}

var chosenColors = map[complex128]color.RGBA{}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Точка (px, py) представляет комплексное значение z.
			img.Set(px, py, z4(z))
		}
	}

	// Создание изображения из стандартного вывода создает битый файл на Win11
	// png.Encode(os.Stdout, img)
	// `go run .\ex3.5.go > image.png`
	f, err := os.Create("mandelbrot.png")
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(f, img); err != nil {
		_ = f.Close()
		log.Fatal(err)
	}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//
//	= z - (z^4 - 1) / (4 * z^3)
//	= z - (z - 1/z^3) / 4
func z4(z complex128) color.Color {
	f := func(z complex128) complex128 {
		return z*z*z*z - 1
	}
	fPrime := func(z complex128) complex128 {
		return (z - 1/(z*z*z)) / 4
	}
	return newton(z, f, fPrime)
}

func round(f float64, digits int) float64 {
	if math.Abs(f) < 0.5 {
		return 0
	}
	pow := math.Pow10(digits)
	return math.Trunc(f*pow+math.Copysign(0.5, f)) / pow
}

func newton(z complex128, f, fPrime Func) color.Color {
	const iterations = 37
	for i := uint8(0); i < iterations; i++ {
		z -= fPrime(z)
		if cmplx.Abs(f(z)) < 1e-6 {
			root := complex(round(real(z), 4), round(imag(z), 4))
			c, ok := chosenColors[root]
			if !ok {
				if len(palettes) == 0 {
					panic("no colors left")
				}
				c = palettes[0]
				palettes = palettes[1:]
				chosenColors[root] = c
			}
			y, cb, cr := color.RGBToYCbCr(c.R, c.G, c.B)
			scale := math.Log(float64(i) / math.Log(iterations))
			y -= uint8(float64(y) * scale)
			return color.YCbCr{y, cb, cr}
		}
	}
	return color.Black
}
