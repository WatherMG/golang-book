/*
	Example 1-3-3

Dup3 выводит количество повторов и строки, которые появляются во входных данных более одного раза.
Программа только список именованных файлов
Аналог команды `uniq` из Unix, ищет повторяющиеся строки
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename) // Возвращает байтовый срез
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") { // Преобразуем байтовый срез в string, чтобы Split мог его разбить по разделителю
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
