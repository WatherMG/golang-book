/*
В этом примере определяется интерфейс StringWriter, который имеет только метод WriteString.
Затем определяется функция writeString, которая принимает аргументы w и s типов io.Writer и string соответственно.
Внутри этой функции используется декларация типа для проверки, соответствует ли динамический тип w интерфейсу StringWriter.
Если это так, то вызывается метод WriteString, чтобы избежать копирования.
В противном случае используется временная копия и вызывается метод Write.

В функции main создаются объекты типов *os.File и *strings.Builder.
Затем используется функция writeString для записи строки в каждый из этих объектов.
Так как тип *os.File имеет метод WriteString, то вызывается этот метод для эффективной записи строки без создания временной копии.
Тип *strings.Builder не имеет метода WriteString, поэтому используется временная копия и вызывается метод Write.

Этот пример показывает, как можно использовать декларации типов для определения новых интерфейсов и проверки типов во
время выполнения для изменения поведения кода в зависимости от типа объекта.
*/

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// StringWriter - интерфейc, который имеет только один метод WriteString.
type StringWriter interface {
	WriteString(string) (n int, err error)
}

// WriteString - функция, которая записывает строку s в w.
// Если w имеет метод WriteString, он вызывается вместо w.Write.
func writeString(w io.Writer, s string) (n int, err error) {
	// Проверка типа с помощью декларации типа: проверяет соответствует ли динамический тип w
	// интерфейсу StringWriter.
	if sw, ok := w.(StringWriter); ok {
		return sw.WriteString(s) // Избегаем копирования.
		// Запись происходит в реализации метода WriteString для конкретного типа. В данном случае в *strings.Builder.
	}
	return w.Write([]byte(s)) // Используем временную копию.
}

func main() {
	// Создаем объекты типов *os.File, *strings.Builder.
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	builder := &strings.Builder{}

	// Записываем строку в каждый из этих объектов с помощью функции writeString.
	if _, err := writeString(file, "Hello, file\n"); err != nil {
		log.Fatal(err)
	}
	if _, err := writeString(builder, "Hello, builder\n"); err != nil {
		log.Fatal(err)
	}

	// Выводим содержимое объекта типа *strings.Builder.
	fmt.Println(builder.String())

	// Закрываем файл.
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
