/*
Exercise 7.2
Напишите функцию CountingWriter с приведенной ниже сигнатурой, которая для данного io.Writer возвращает новый Writer,
являющийся оболочкой исходного, и указатель на переменную int64, которая в любой момент содержит количество байтов,
записанных в новый Writer.

func CountingWriter(w io.Writer) (io.Writer, *int64)
*/

package main

import (
	"fmt"
	"io"
)

type ByteCounter struct {
	w       io.Writer // Интерфейс, в который записываются данные
	written int64     // Количество записанных байтов
}

func main() {
	w, c := CountingWriter(io.Discard) // Создаем новый io.Writer, с помощью CountingWriter
	fmt.Fprintf(w, "Hello world!")     // Записываем строку в этот io.Writer
	fmt.Println(*c)                    // Выводим количество записанных байтов
}

// Write реализация интерфейса io.Writer. Записывает переданный в него срез байтов p в w и увеличивает
// значение written на количество записанных байт. Возвращает количество записанных байт и ошибку.
func (b *ByteCounter) Write(p []byte) (n int, err error) {
	n, err = b.w.Write(p)
	b.written += int64(n)
	return n, err
}

// CountingWriter создает новый объект ByteCounter, используя переданный интерфейс io.Writer
// как исходный параметр, и возвращает его, а так же указатель на поле written.
// С помощью него мы можем получить количество байтов, записанных в переданный интерфейс io.Writer
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &ByteCounter{w, 0}
	return c, &c.written
}
