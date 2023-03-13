/*
Example.5.10
Findlinks3 просматривает веб-страницы, начиная с URL-адресов, указанных в командной строке.
*/

package main

import (
	"fmt"
	"log"
	"os"

	"GolangBook/chapter5/lesson6/links"
)

// breadthFirst вызывает f для каждого элемента в worklist.
// Все элементы, возвращаемые f, добавляются в worklist.
// f вызывается для каждого элемента не более одного раза
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
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

func main() {
	// Поиск в ширину, начиная с аргумента командной строки
	breadthFirst(crawl, os.Args[1:])
}
