package main

import "fmt"

// Animal У нас есть интерфейс, который определяет метод Speak.
// и есть два типа Dog, Cat, которые реализуют интерфейс Animal
type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "bark"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "meow"
}

func main() {
	// Создаем переменную типа Animal и присваиваем ей значение типа Dog или Cat
	var a Animal
	a = Dog{}

	// Используя декларацию типа (type assertion), проверяем, является ли значение a Dog или Cat
	// Если декларация типа успешна (т.е. значение a имеет тип Dog или Cat), то мы
	// извлекаем значение из интерфейса и используем его. Если декларация типа не успешна,
	// то мы выводим сообщение об ошибке.
	if d, ok := a.(Dog); ok {
		fmt.Println("it's a dog:", "say:", d.Speak())
	} else if c, ok := a.(Cat); ok {
		fmt.Println("it's a cat:", "say:", c.Speak())
	} else {
		fmt.Println("a is not a dog or a cat")
	}
}
