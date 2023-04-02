package main

import (
	"fmt"
	"math"
)

/*
Декларация типа (type assertion) в Go используется не только для проверки типа значения интерфейса,
но и для извлечения значения из интерфейса. Это позволяет получить доступ к методам и полям конкретного типа,
которые не определены в интерфейсе.
*/

// Shape Создадим интерфейс, который определяет метод Area
type Shape interface {
	Area() float64
}

// Rectangle создадим тип, который реализует этот интерфейс
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) Diameter() float64 {
	return 2 * c.radius
}

func main() {
	// Создадим переменную типа Shape и присвоим ей значение типа Rectangle
	var s Shape
	s = Rectangle{3, 4} // s имеет тип интерфейса Shape с динамическим типом Rectangle и его значением

	// s.Area() у s мы имеем доступ только к методам типа

	// Мы можем использовать декларацию типа, чтобы извлечь значение Rectangle
	// из интерфейса Shape и получить доступ к полям Width и Height, которые не определены в интерфейсе
	r := s.(Rectangle) // r имеет тип Rectangle и доступ к его полям и методам
	fmt.Println(r.Width, r.Height)

	s = Circle{2.5}
	c := s.(Circle) // извлекаем значение Circle из интерфейса Shape.
	fmt.Printf("circle: Radius=%.2f, Area=%.2f, Diameter=%.2f", c.radius, c.Area(), c.Diameter())
}
