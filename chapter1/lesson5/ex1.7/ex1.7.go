/*
Exercise 1-7
Fetch выводит ответ на запрос по-заданному URL.
"Аналог curl"

Вызов функции io.Copy(dst, src) выполняет чтение src и запись в dst. Воспользуйтесь ею вместо ioutil. ReadAll для
копирования тела ответа в поток os. Stdout без необходимости выделения достаточно большого для хранения
всего ответа буфера. Не забудьте проверить, не произошла ли ошибка при вызове io. Сору
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		response, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		nbytes, err := io.Copy(os.Stdout, response.Body)
		_ = response.Body.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("\n%d", nbytes) // Выводим размер ответа
	}
}
