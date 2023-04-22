/*
Exercise 8.11
Следуя подходу mirroredQuery из раздела 8.4.4, реализуйте вариант
программы fetch, который параллельно запрашивает несколько URL. Как только
получен первый ответ, прочие запросы отменяются.
*/

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Printf("empty URLs\n")
		return
	}

	// Каналы отмены и ответов
	cancel := make(chan struct{})
	responses := make(chan string, len(os.Args[1:]))

	for _, url := range os.Args[1:] {
		// Для всех адресов создаем горутины, которые делают запрос и отправляют адрес в канал
		go func(url string) {
			responses <- fetch(url, cancel)
		}(url)
	}
	// Получаем адрес, который первый попал в канал и закрываем их
	resp := <-responses
	close(cancel)
	fmt.Println(resp)
}

func fetch(url string, cancelled <-chan struct{}) string {
	cxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	req, err := http.NewRequestWithContext(cxt, http.MethodGet, url, nil)

	if err != nil {
		return err.Error()
	}
	go func() {
		select {
		case <-cancelled:
			cancel()
		case <-cxt.Done():
			fmt.Println()
		}

	}()
	t := time.Now()
	resp, err := http.DefaultClient.Do(req)
	elapsed := time.Since(t)

	if err != nil {
		return err.Error()
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// Такое большое время ответа от сервера, по сравнению с утилитой ping зависит от используемого протокола.
		// В ping используется ICMP, а в этом приложении HTTP/1
		return fmt.Sprintf("%s: %s", resp.Request.URL.Host, elapsed)
	}
	return resp.Request.URL.Host + " " + resp.Status
}
