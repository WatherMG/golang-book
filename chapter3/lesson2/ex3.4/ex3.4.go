/*
Exercise 3.4
Surface вычисляет SVG-представление трехмерного графика функции.
Следуя подходу, использованному в примере с фигурами Лиссажу из раздела 1.7, создайте веб-сервер,
который вычисляет поверхности и возвращает клиенту SVG-данные.
Сервер должен использовать в ответе заголовок ContentType наподобие следующего:
`w.Header().Set("ContentType", "image/svg+xml")`
*/

package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

const (
	cells   = 100         // Количество ячеек сетки
	xyrange = 30          // Диапазон осей (-xyrange..+xyrange)
	angle   = math.Pi / 6 // Углы осей x, y (=30°)
)

var width, height float64 // Размер канвы в пикселях
var peakColor, lowlandColor string
var formType string                                 // Цвет пиков и низин
var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

// exercise 3.4
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/saddle", handler)
	http.HandleFunc("/climb", handler)
	http.HandleFunc("/drop", handler)
	// fun := r.URL.Query().Get("fun")
	log.Fatal(http.ListenAndServe("192.168.0.199:8000", nil))
}

// handler получает форму объекта, получает параметры и запускает функцию получения объекта
func handler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.RequestURI(), "drop") {
		formType = "drop"
	} else if strings.Contains(r.URL.RequestURI(), "saddle") {
		formType = "saddle"
	} else {
		formType = "climb"
	}
	getParams(r)
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w)
}

func surface(out io.Writer) {
	_, _ = fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg'> "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", int(width), int(height))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			// exercise 3.1
			ax, ay, ct, err := corner(i+1, j)
			if err != nil {
				continue
			}
			bx, by, ct1, err := corner(i, j)
			if err != nil {
				continue
			}
			cx, cy, ct2, err := corner(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, ct3, err := corner(i+1, j+1)
			if err != nil {
				continue
			}

			// exercise 3.3
			var color string

			switch {
			case ct == 1 || ct1 == 1 || ct2 == 1 || ct3 == 1:
				color = peakColor
			case ct == 2 || ct1 == 2 || ct2 == 2 || ct3 == 2:
				color = lowlandColor
			default:
				color = "#00ff00"
			}
			_, _ = fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
				"fill='#222222' stroke='%s' stroke-width='0.4'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	_, _ = fmt.Fprintf(out, "</svg>")
}

// getParams получает параметры из URL
func getParams(r *http.Request) {
	width = 600
	if canvasWidth, err := strconv.ParseFloat(r.URL.Query().Get("w"), 64); err == nil {
		width = canvasWidth
	}

	height = 320
	if canvasHeight, err := strconv.ParseFloat(r.URL.Query().Get("h"), 64); err == nil {
		height = canvasHeight
	}

	peakColor = "#0000ff"
	if pc := r.URL.Query().Get("pc"); pc != "" {
		peakColor = pc
	}

	lowlandColor = "#ff0000"
	if lc := r.URL.Query().Get("lc"); lc != "" {
		lowlandColor = lc
	}
}

// calcSize Возвращает количество пикселей в x или y и количество пикселей в z
func calcSize(w, h float64) (xyscale, zscale float64) {
	return w / 2 / xyrange,
		h * 0.4
}

func corner(i, j int) (float64, float64, int, error) {
	// Ищем угловую точку (x,y) ячейки (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Вычисляем высоту поверхности z
	z, ct := f(x, y)
	// exercise 3.1
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, 0, fmt.Errorf("invalid value")
	}
	// exercise 3.4
	xyscale, zscale := calcSize(width, height)

	// Изометрически проецируем (x, y, z) на двумерную канву SVG (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, ct, nil

}

func f(x, y float64) (float64, int) {
	var z float64
	switch formType {
	case "drop":
		z = dropForm(x, y)
	case "climb":
		z = climbForm(x, y)
	case "saddle":
		z = saddleForm(x, y)
	}
	ct := 0
	// exercise 3.3 получаем пики и низины
	if z < 0 {
		ct = 2
	} else {
		ct = 1
	}
	return z, ct
}

func climbForm(x, y float64) float64 {
	return (math.Sin(x) / x) * (math.Sin(y) / y)
}

func saddleForm(x, y float64) float64 {
	return math.Pow(x, 2)/math.Pow(25, 2) - math.Pow(y, 2)/math.Pow(17, 2)
}

func dropForm(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
