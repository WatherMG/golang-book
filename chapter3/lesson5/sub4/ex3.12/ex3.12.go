/*
Exercise 3.12
Напишите функцию, которая сообщает, являются ли две строки анаграммами одна другой,
т.е. состоят ли они из одних и тех же букв в другом порядке.

*/

package anagram

import (
	"strings"
	"unicode"
)

func isAnagram(f, s string) bool {
	// Создаем хэш таблицы для обоих строк
	firstSeq := makeRuneMap(f)
	secondSeq := makeRuneMap(s)

	// Проверяем совпадает ли количество букв первого слова с количеством букв второго слова и наоборот
	if isMatchLetters(firstSeq, secondSeq) && isMatchLetters(secondSeq, firstSeq) {
		return true
	}
	return false
}

// makeRuneMap в цикле составляет словарь с количеством букв в переданной строке
func makeRuneMap(s string) map[rune]uint8 {
	seq := make(map[rune]uint8)
	s = strings.ToLower(s)
	for _, v := range s {
		if unicode.IsLetter(v) {
			seq[v]++
		}
	}
	return seq
}

// isMatchLetters проверяет количество букв одной хеш-таблицы с другой
func isMatchLetters(f, s map[rune]uint8) bool {
	for k, v := range f {
		if s[k] != v {
			return false
		}
	}
	return true
}
