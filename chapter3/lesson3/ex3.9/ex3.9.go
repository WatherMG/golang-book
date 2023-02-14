/*
Exercise 3.5
Mandelbrot создает PNG-изображение фрактала Мандельброта.
Напишите программу веб-сервера, который визуализирует фракталы и выводит данные изображения клиенту.
Позвольте клиенту указывать значения х, у и масштабирования в качестве параметров HTTP-запроса.
*/

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
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
var fractalType float64

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("192.168.0.199:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		_, _ = fmt.Fprintf(w, "Error parsing form: %v", err)
	}
	x := getParamFromURL(r.Form["x"], 0)
	y := getParamFromURL(r.Form["y"], 0)
	scale := getParamFromURL(r.Form["s"], 0)
	fractalType = getParamFromURL(r.Form["ft"], -1)
	render(w, x, y, scale)
}

func getParamFromURL(params []string, def float64) float64 {
	if len(params) == 0 {
		return def
	}
	value, err := strconv.ParseFloat(params[0], 64)
	if err != nil {
		return def
	}
	return value
}

func render(w http.ResponseWriter, x, y, zoom float64) {
	const (
		width, height = 1024, 1024
	)

	scale := math.Exp2(1 - zoom)
	xmin, xmax := x-scale, x+scale
	ymin, ymax := y-scale, y+scale

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Точка (px, py) представляет комплексное значение z.
			switch fractalType {
			case 0:
				img.Set(px, py, newton(z))
			case 1:
				img.Set(px, py, acos(z))
			case 2:
				img.Set(px, py, sqrt(z))
			default:
				img.Set(px, py, mandelbrot(z))

			}

		}
	}
	err := png.Encode(w, img)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	fractalType = 4

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
	fractalType = 1
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{Y: 192, Cb: blue, Cr: red}
}

func sqrt(z complex128) color.Color {
	fractalType = 2
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
	fractalType = 0
	const iterations = 37
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return palettes[i%16]
		}
	}
	return color.Black
}
