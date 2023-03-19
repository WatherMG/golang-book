package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	distanceFromP := p.Distance        // Значение метод
	fmt.Println(distanceFromP(q))      // 5
	var origin Point                   // {0, 0}
	fmt.Println(distanceFromP(origin)) // 2.23606797749979, sqrt(5)

	scaleP := p.ScaleBy // Значение-метод
	scaleP(2)           // p становится (2, 4)
	scaleP(3)           // затем (6, 12)
	scaleP(10)          // затем (60, 120)

	// Тут будет выведено 2.23... т.к. изменения внесенные через scaleP не вносят изменения в переменную distanceFromP.
	// Она все еще содержит оригинальную точку p = {1, 2}
	fmt.Println(distanceFromP(origin)) // 2.23606797749979, sqrt(5)
}
