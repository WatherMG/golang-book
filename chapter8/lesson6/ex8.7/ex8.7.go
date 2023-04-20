/*
Exercise 8.7
Напишите параллельную программу, которая создает локальное зеркало
веб-сайта, загружая все доступные страницы и записывая их в каталог на локальном
диске. Выбираться должны только страницы в пределах исходного домена (например,
golang.org). URL в страницах зеркала должны при необходимости быть изменены
таким образом, чтобы они ссылались на зеркальную страницу, а не на оригинал.
*/

package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"GolangBook/chapter5/lesson6/links"
)

var (
	wg   = &sync.WaitGroup{}
	base = flag.String("u", "https://remontpc82.ru/", "base url to crawl")
)

func main() {
	flag.Parse()

	for _, url := range crawl(*base) {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			download(*base, url)
		}(url)
	}
	done := make(chan struct{})

	go func() {
		wg.Wait()
		done <- struct{}{}
	}()
	<-done
}

func download(base, url string) {
	if !strings.HasPrefix(url, base) || strings.Contains(url, "#") {
		return
	}

	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	dir := strings.TrimPrefix(strings.TrimPrefix(url, "http://"), "https://")
	if err := os.MkdirAll(dir, 0755); err != nil && !os.IsExist(err) {
		log.Println(err)
	}
	ext := "index.html"
	filename := dir + ext
	f, err := os.Create(filename)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	_, err = io.Copy(f, res.Body)
	if err != nil {
		log.Println(err)
	}
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Println(err)
	}
	return list
}
