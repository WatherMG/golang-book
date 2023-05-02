/*
Exercise 11.1
Напишите тесты для программы charcount из раздела 4.3.
Example 4.4
Charcount вычисляет количество символов Unicode.
*/

package charcount

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func GetCharCount(r io.Reader) [5]int {
	counts := make(map[rune]int)    // Количество символов Unicode
	var utflen [utf8.UTFMax + 1]int //  Количество длин кодировок UTF-8
	invalid := 0                    // Количество некорректных символов UTF-8

	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune() // Возвращает руну, байты, ошибки
		if errors.Is(err, io.EOF) {
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
		counts[r]++
		utflen[n]++
	}
	return utflen
}
