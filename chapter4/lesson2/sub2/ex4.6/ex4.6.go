/*
Exercise 4.6

Напишите функцию, которая без выделения дополнительной памяти преобразует последовательности смежных пробельных
символов Unicode (см. Unicode.IsSpace) в срезе []byte в кодировке UTF-8 в один пробел ASCII.
*/

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	data := []byte("This is  test  text  ")
	fmt.Printf("%s\n", convert(data))
}

func convert(data []byte) []byte {
	for i := 0; i < len(data); {
		first, size := utf8.DecodeRune(data[i:])
		if unicode.IsSpace(first) {
			second, _ := utf8.DecodeRune(data[i+size:])
			if unicode.IsSpace(second) {
				copy(data[i:], data[i+size:])
				data = data[:len(data)-size]
				continue
			}
		}
		i += size
	}
	return data
}
