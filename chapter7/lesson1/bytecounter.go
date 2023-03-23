/*
Example 7.1
Bytecounter подсчитывает количество записанных в него байт и накапливает длину выводимого результата.
*/

package main

import "fmt"

type ByteCounter int

// Write. Для соответствия типа len(p) типу *c в операторе += необходимо выполнить преобразование типа.
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // Преобразование int в ByteCounter
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // 5, = len("hello")
	c = 0          // сброс счетчика
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // 12, len=("hello, Dolly")

}
