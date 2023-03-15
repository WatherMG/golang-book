/*
Exercise 5.18
Перепишите, не изменяя ее поведение, функцию fetch так, чтобы она использовала defer для закрытия записываемого файла.

Решение соответствует информации из последнего абзаца:
если и io.Copy, и f.Close завершаются неудачно, следует предпочесть отчет об ошибке io.Copy, поскольку она произошла
первой и, скорее всего, сообщит главную причину неприятностей.
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	for _, u := range os.Args[1:] {
		local, n, err := fetch(u)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", u, err)
			continue
		}
		fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", u, local, n)
	}
}

// Fetch загружает URL и возвращает имя и длину локального файла.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)

	// Закрываем файл. Если ошибка есть в Copy - возвращаем ее
	// Если ошибка есть в f.Close и нет в Copy - возвращаем f.Close
	// Если ошибка есть и в f.Close и в Copy - возвращаем Copy
	defer func() {
		if closeErr := f.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	return local, n, err
}
