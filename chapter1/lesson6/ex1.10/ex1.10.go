/*
Exercise 1.10
Fetchall выполняет параллельную выборку URL и сообщает о затраченном времени и размере ответа для каждого из них.
"Аналог curl"
Найдите веб-сайт, который содержит большое количество данных. Исследуйте работу кеширования путем двукратного
запуска fetchall и сравнения времени запросов. Получаете ли вы каждый раз одно и то же содержимое?
Измените fetchall так, чтобы вывод осуществлялся в файл и чтобы затем можно было его изучить.
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
			url = "https://" + url
		}
		go fetch1(url, ch) // Запуск горутины
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.8fs elapsed\n", time.Since(start).Seconds())
}

func fetch1(url string, ch chan<- string) {
	start := time.Now()

	response, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // Отправка ошибки в канал `ch`
		return
	}

	f, err := os.Create(strings.Replace(url, "://", "_", 1) + "-dump.html")
	if err != nil {
		ch <- err.Error()
	}

	nbytes, err := io.Copy(f, response.Body)
	_ = response.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.8fs %7d %s", secs, nbytes, url)
}
