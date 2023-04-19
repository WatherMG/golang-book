/*
Crawl1 просматривает веб-ссылки, начиная с аргументов командной строки.
Эта версия быстро исчерпывает доступные файловые дескрипторы из-за
чрезмерного количества одновременных вызовов links.Extract. Кроме того,
она никогда не завершается, поскольку рабочий список никогда не закрывается.
*/
package main

import (
	"fmt"
	"log"
	"os"

	"GolangBook/chapter5/lesson6/links"
)

func main() {
	worklist := make(chan []string)

	// Запуск с аргументами командной строки
	go func() { worklist <- os.Args[1:] }()

	// Параллельное сканирование
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
