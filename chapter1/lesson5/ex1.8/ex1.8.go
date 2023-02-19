/*
Exercise 1-8
Fetch выводит ответ на запрос по-заданному URL.
"Аналог curl"

Измените программу fetch так, чтобы к каждому аргументу URL автоматически добавлялся префикс http://
в случае отсутствия в нем такового. Можете воспользоваться функцией strings. HasPrefix.
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		response, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		body, err := io.ReadAll(response.Body)
		_ = response.Body.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", body)
	}
}
