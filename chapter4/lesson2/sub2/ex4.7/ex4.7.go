/*
Exercise 4.7

Перепишите функцию reverse так, чтобы она без выделения дополнительной памяти обращала последовательность
символов среза []byte, который представляет строку в кодировке UTF-8.
Сможете ли вы обойтись без выделения новой памяти?
*/

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	a := "一 二 三"
	fmt.Printf("%s\n", a)

	fmt.Printf("%s\n", revUTF8([]byte(a)))
}

func reverse(s []byte) {
	size := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[size-i] = s[size-i], s[i]
	}
}

func revUTF8(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverse(b[i : i+size])
		i += size
	}
	reverse(b)
	return b
}
