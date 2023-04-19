/*
Crawl3 просматривает веб-ссылки, начиная с аргументов командной строки.
Эта версия использует ограниченный параллелизм. Для простоты она не
рассматривает проблему завершения.
*/
package main

import (
	"fmt"
	"log"
	"os"

	"GolangBook/chapter5/lesson6/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
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

	// Создание 20 горутин сканирования для выборки всех непросмотренных ссылок.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
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
