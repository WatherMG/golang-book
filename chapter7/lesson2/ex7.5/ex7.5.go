/*
Exercise 7.5
Функция LimitReader из пакета io принимает переменную r типа io.Reader и количество байтов n и возвращает
другой объект Reader, который читает из r, но после чтения n сообщает о достижении конца файла. Реализуйте его.

func LimitReader(r io.Reader, n int64) io.Reader
*/

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type LimitedReader struct {
	r        io.Reader
	n, limit int64
}

func (r *LimitedReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p[:r.limit])
	r.n += int64(n)
	if r.n >= r.limit {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, limit int64) io.Reader {
	return &LimitedReader{r: r, limit: limit}
}

func main() {
	lr := LimitReader(strings.NewReader("123456789"), 3)
	b, err := io.ReadAll(lr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
	}
	fmt.Printf("%s\n", b)
}
