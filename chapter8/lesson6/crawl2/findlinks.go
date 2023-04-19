/*
Crawl2 просматривает веб-ссылки, начиная с аргументов командной строки.
Эта версия использует буферизованный канал в качестве счетного семафора
для ограничения количества одновременных вызовов links.Extract.
*/
package main

import (
	"fmt"
	"log"
	"os"

	"GolangBook/chapter5/lesson6/links"
)

// tokens представляет собой подсчитывающий семафор, используемый
// для ограничения количества параллельных запросов величиной 20.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // Захват токена
	list, err := links.Extract(url)
	<-tokens // Освобождение токена
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int // Количество ожидающих отправки в рабочий список

	// Запуск с аргументами командно строки.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Параллельное сканирование веб.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
