package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 8      // Эквивалентно записи w.Circle.Point.X = 8
	w.Y = 8      // Эквивалентно записи w.Circle.Point.Y = 8
	w.Radius = 5 // Эквивалентно записи w.Circle.Radius = 5
	w.Spokes = 20

	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	w = Wheel{
		Circle: Circle{
			Point:  Point{8, 8},
			Radius: 5,
		},
		Spokes: 20, // Примечание: завершающая запятая здесь (и после Radius) необходима
	}

	fmt.Printf("%#v\n", w) // main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20}

	w.X = 42
	fmt.Printf("%#v\n", w) // main.Wheel{Circle:main.Circle{Point:main.Point{X:42, Y:8}, Radius:5}, Spokes:20}
}
