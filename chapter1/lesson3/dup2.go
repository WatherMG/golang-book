/*
	Example 1-3-2

Dup2 выводит количество повторов и строки, которые появляются во входных данных более одного раза.
Программа читает стандартный ввод или список именованных файлов
Аналог команды `uniq` из Unix, ищет повторяющиеся строки
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts) // читаем ввод в консоль
	} else {
		for _, arg := range files { // получаем список файлов
			f, err := os.Open(arg)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue

			}
			countLines(f, counts)
			_ = f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// Игнорируем потенциальные ошибки из input.Err()
}
