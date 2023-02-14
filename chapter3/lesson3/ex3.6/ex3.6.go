/*
Exercise 3.6
Mandelbrot создает PNG-изображение фрактала Мандельброта.
Реализуйте полноцветное множество Мандельброта с использованием функции image.NewRGBA и типа color.RGB А или color.YCbCr
Супервыборка (supersampling) — это способ уменьшить эффект пикселизации путем вычисления
значений цвета в нескольких точках в пределах каждого пикселя и их усреднения.
Проще всего разделить каждый пиксель на четыре “подпикселя”. Реализуйте описанный метод.
*/

package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
)

var palettes = [...]color.RGBA{
	{66, 30, 15, 255},
	{25, 7, 26, 255},
	{9, 1, 47, 255},
	{4, 4, 73, 255},
	{0, 7, 100, 255},
	{12, 44, 138, 255},
	{24, 82, 177, 255},
	{57, 125, 209, 255},
	{134, 181, 229, 255},
	{211, 236, 248, 255},
	{241, 233, 191, 255},
	{248, 201, 95, 255},
	{255, 170, 0, 255},
	{204, 128, 0, 255},
	{153, 87, 0, 255},
	{106, 52, 3, 255},
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		espX                   = (xmax - xmin) / width
		espY                   = (ymax - ymin) / height
	)

	offX := []float64{-espX, espX}
	offY := []float64{-espY, espY}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// Supersampling
			subPixels := make([]color.Color, 0)
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					z := complex(x+offX[i], y+offY[j])
					subPixels = append(subPixels, mandelbrot(z))
				}
			}
			// Точка (px, py) представляет комплексное значение z.
			img.Set(px, py, avg(subPixels))
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

func avg(colors []color.Color) color.Color {
	var r, g, b, a uint16
	n := len(colors)
	for _, c := range colors {
		colorR, colorG, colorB, colorA := c.RGBA()
		r += uint16(colorR / uint32(n))
		g += uint16(colorG / uint32(n))
		b += uint16(colorB / uint32(n))
		a += uint16(colorA / uint32(n))
	}
	return color.RGBA64{r, g, b, a}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200

	var v complex128
	for n := uint16(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palettes[n%16]
		}
	}
	return color.Black
}

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{Y: 128, Cb: blue, Cr: red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//
//	= z - (z^4 - 1) / (4 * z^3)
//	= z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
