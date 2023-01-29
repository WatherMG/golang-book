/*
Exercise 1-5-3
Fetch выводит ответ на запрос по-заданному URL.
"Аналог curl"

Измените программу fetch так, чтобы она выводила код состояния HTTP, содержащийся в resp. Status.
*/

package main

import (
	"fmt"
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
		fmt.Printf("%s: %s\n", url, response.Status)
	}
}
