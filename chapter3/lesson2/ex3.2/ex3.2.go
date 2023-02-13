/*
Exercise 3.2
Surface вычисляет SVG-представление трехмерного графика функции.
Поэкспериментируйте с визуализациями других функций из пакета math.
Сможете ли вы получить изображения наподобие коробки для яиц, седла или холма?
*/

package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // Размер канвы в пикселях
	cells         = 100                 // Количество ячеек сетки
	xyrange       = 30                  // Диапазон осей (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // Пикселей в единице x или y
	zscale        = height * 0.4        // Пикселей в единице z
	angle         = math.Pi / 6         // Углы осей x, y (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg'> "+
		"style='stroke: gray; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			// exercise 3.1
			ax, ay, err := corner(i+1, j)
			if err != nil {
				continue
			}
			bx, by, err := corner(i, j)
			if err != nil {
				continue
			}
			cx, cy, err := corner(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, err := corner(i+1, j+1)
			if err != nil {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, error) {
	// Ищем угловую точку (x,y) ячейки (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Вычисляем высоту поверхности z
	z := f(x, y)
	// exercise 3.1
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, fmt.Errorf("invalid value")
	}
	// Изометрически проецируем (x, y, z) на двумерную канву SVG (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil

}

func f(x, y float64) float64 {
	return eggform(x, y)
}

func climb(x, y float64) float64 {
	return (math.Sin(x) / x) * (math.Sin(y) / y)
}

func eggform(x, y float64) float64 {
	return math.Pow(2, math.Sin(x)) * math.Pow(2, math.Sin(y)) / 12
}

func saddle(x, y float64) float64 {
	return math.Pow(x, 2)/math.Pow(25, 2) - math.Pow(y, 2)/math.Pow(17, 2)
}
