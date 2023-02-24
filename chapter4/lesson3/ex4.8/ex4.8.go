/*
Exercise 4.8
Измените сhaгcount так, чтобы программа подсчитывала количество букв, цифр и прочих категорий Unicode
с использованием функций наподобие Unicode.IsLetter
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int) // Количество символов Unicode
	categories := make(map[string]int)
	var utflen [utf8.UTFMax + 1]int //  Количество длин кодировок UTF-8
	invalid := 0                    // Количество некорректных символов UTF-8

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // Возвращает руну, байты, ошибки
		if err == io.EOF {
			break
		}
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		for catName, rangeTable := range unicode.Properties {
			if unicode.In(r, rangeTable) {
				categories[catName]++
			}
		}
		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Printf("\n%-34s count\n", "category")
	for c, n := range categories {
		fmt.Printf("%-34s %d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
