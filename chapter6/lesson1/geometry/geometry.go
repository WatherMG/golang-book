/*
Example 6.1
Пакет geometry определяет простые типы для геометрии плоскости
*/

package geometry

import (
	"math"
)

type Point struct {
	X, Y float64
}

// Path - путь из точек, соединенных прямолинейными отрезками.
type Path []Point

// традиционная функция
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// То же, но как метод типа Point
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance возвращает длину пути
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
