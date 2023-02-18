/*
Exercise 3.11
Усовершенствуйте функцию comma так, чтобы она корректно работала с числами с плавающей точкой и необязательным знаком.
*/

package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := 3220.99
	s := strconv.FormatFloat(a, 'f', 3, 64)
	fmt.Println(comma(s))
}

// comma разбивает строку символом ',' каждые 3 символа строки
func comma(s string) string {
	// Используем указатель на буфер, чтобы не пересоздавать его, в случае нескольких вызовов
	b := &bytes.Buffer{}
	// Определяем число положительное или отрицательное и устанавливаем начальный индекс строки
	signIndex := 0
	if s[0] == '-' || s[0] == '+' {
		b.WriteByte(s[0])
		signIndex = 1
	}
	// Определяем индекс точки
	commaIndex := strings.Index(s, ".")
	if commaIndex == -1 {
		commaIndex = len(s)
	}
	// Получаем целую часть
	text := s[signIndex:commaIndex]
	// Определяем сколько символов должно быть в первой группе
	n := len(text) % 3
	if n > 0 {
		// Записываем первую группу
		b.Write([]byte(text[:n]))
		if len(text) > n {
			b.WriteString(",")
		}
	}
	// В цикле расставляем запятую, перед каждым 3-м символом из целой части и добавляем в буфер символы
	for i, c := range text[n:] {
		if i%3 == 0 && i != 0 {
			b.WriteRune(',')
		}
		b.WriteRune(c)
	}
	// Добавляем дробную часть
	b.WriteString(s[commaIndex:])

	return b.String()
}
