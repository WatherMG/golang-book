/*
Example 5.15
Fetch сохраняет содержимое URL в локальный файл.
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
	// Закрытие файла; если есть ошибка Copy - возвращаем ее
	if closeErr := f.Close(); err != nil {
		err = closeErr
	}
	return local, n, err
}
