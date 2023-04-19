/*
Exercise 8.6
Добавьте ограничение по глубине в параллельный сканер. Иначе
говоря, если пользователь устанавливает -depth=3, то выбираются только те URL,
которые достижимы через цепочку не более чем из трех ссылок.
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"GolangBook/chapter5/lesson6/links"
)

var maxDepth int

type work struct {
	url   string
	depth int
}

// tokens представляет собой подсчитывающий семафор, используемый
// для ограничения количества параллельных запросов величиной 20.
var tokens = make(chan struct{}, 20)

func crawl(w work) []work {
	fmt.Printf("Depth: %d, URL: %s\n", w.depth, w.url)

	if w.depth >= maxDepth {
		return nil
	}

	tokens <- struct{}{} // Захват токена
	list, err := links.Extract(w.url)
	<-tokens // Освобождение токена
	if err != nil {
		log.Print(err)
	}

	works := make([]work, len(list))

	for _, url := range list {
		works = append(works, work{url: url, depth: w.depth + 1})
	}
	return works
}

func main() {
	flag.IntVar(&maxDepth, "d", 3, "max depth")
	flag.Parse()
	worklist := make(chan []work)
	var n int // Количество ожидающих отправки в рабочий список

	// Запуск с аргументами командно строки.
	n++
	go func() {
		var works []work
		for _, url := range os.Args[3:] {
			works = append(works, work{url, 1})
		}
		worklist <- works
	}()

	// Параллельное сканирование веб.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		works := <-worklist
		for _, w := range works {
			if !seen[w.url] {
				seen[w.url] = true
				n++
				go func(w work) {
					worklist <- crawl(w)
				}(w)
			}
		}
	}
}
