/*
Exercise 8.10
Запросы HTTP могут быть отменены с помощью закрытия
необязательного канала Cancel в структуре http.Request. Измените веб-сканер из
раздела 8.6 так, чтобы он поддерживал отмену. Указание. Функция http.Get не
позволяет настроить Request. Вместо этого создайте запрос с использованием
http.NewRequest, установите его поле Cancel и выполните запрос с помощью вызова
http.DefaultClient.Do(req).
*/

package main

import (
	"fmt"
	"log"
	"os"

	"GolangBook/chapter8/lesson9/ex8.10/links"
)

func crawl(url string, cancelled <-chan struct{}) []string {
	fmt.Println(url)
	list, err := links.Extract(url, cancelled)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)  //  Список URL, могут быть дубли.
	unseenLinks := make(chan string) // Удаление дублей.

	// Добавление в список аргументов командной строки
	go func() { worklist <- os.Args[1:] }()

	cancelled := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(cancelled)
	}()

	// Создание 20 горутин сканирования для выборки всех непросмотренных ссылок.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link, cancelled)
				go func() { worklist <- foundLinks }()

			}
		}()
	}

	// Главная горутина удаляет дубликаты из списка и отправляет непросмотренные ссылки сканерам.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
