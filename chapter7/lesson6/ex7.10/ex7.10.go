/*
Exercise 7.10
Тип sort.Interface можно адаптировать для других применений. Напишите функцию IsPalindrome(s sort.Interface) bool,
которая сообщает, является ли последовательность s палиндромом (другими словами, что обращение последовательности не изменяет ее).
Считайте, что элементы с индексами i и j равны, если !s.Less(i, j)&&!s.Less(j, i).

Проблема:
Столкнулся с некорректной работой программы, было связанно с тем, что:
Для русских символов нужно использовать тип []rune. Так как при использовании []byte
русские символы используют несколько байт в UFT-8.
*/

package main

import (
	"fmt"
	"sort"
)

type word []rune

func (w word) Len() int {
	return len(w)
}

func (w word) Less(i, j int) bool {
	return w[i] < w[j]
}

func (w word) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func equal(i, j int, s sort.Interface) bool {
	return !s.Less(i, j) && !s.Less(j, i)
}

func IsPalindrome(s sort.Interface) bool {
	max := s.Len() - 1
	for i := 0; i < s.Len()/2; i++ {
		if !equal(i, max-i, s) {
			return false
		}
	}
	return true
}

func main() {
	s := word("шалаш")
	fmt.Println(IsPalindrome(s))
}
