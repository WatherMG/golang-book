/*
Exercise 4.12
Популярный веб-ресурс с комиксами xkcd имеет интерфейс JSON. Например, запрос https://xkcd.com/571/info.0.json
возвращает детальное описание комикса 571, одного из многочисленных фаворитов сайта.
Загрузите каждый URL (по одному разу!) и постройте автономный список комиксов. Напишите программу xkcd, которая,
используя этот список, будет выводить URL и описание каждого комикса, соответствующего условию поиска,
заданному в командной строке.

The popular web comic xkcd has a JSON interface. For example, a request to https://xkcd.com/571/info.0.json
produces a detailed description of comic 571, one of many favorites.
Download each URL (once!) and build an offline index. Write a tool xkcd that, using this index,
prints the URL and transcript of each comic that matches a search term provided on the command line
*/

package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"GolangBook/chapter4/lesson5/ex4.12/xkcd"
)

func init() {
	flag.Parse()
}

var (
	f = flag.Bool("f", false, "use to fetch comics")
	n = flag.Int("n", 10, "number of comics")
)

// fetch: go run main.go -f -n=100
// search: go run main.go keywords
func main() {
	if *f {
		if *n > xkcd.MaxNum {
			log.Fatalf("%d can't be bigger than %d", *n, xkcd.MaxNum)
		}
		fetchComics(*n, 16)
	} else {
		searchComics(flag.Args())
	}
}

func searchComics(keywords []string) {
	if len(keywords) == 0 {
		fmt.Println("You don't specify keywords")
		os.Exit(1)
	}
	f, err := os.Open("comics.json")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	data := bufio.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}

	f.Seek(0, 0)

	index := xkcd.New(0, 10)
	if err := json.NewDecoder(data).Decode(&index); err != nil {
		log.Fatal(err)
	}

	res := xkcd.Search(index, keywords)
	for _, c := range res {
		if c != nil {
			fmt.Println(c)
		}
	}

}

func fetchComics(n, maxWorkers int) {
	index := xkcd.New(n, n)
	sem := make(chan struct{}, maxWorkers)
	wg := sync.WaitGroup{}

	for i := xkcd.MinNum; i < n+1; i++ {
		sem <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer func() {
				<-sem
				wg.Done()
			}()

			p, err := xkcd.Get(i)
			if err != nil {
				log.Printf("Failed to fetch comic #%d: %v\n", i, err)
				return
			}
			index.Comics[i-1] = p
		}(i)
	}
	wg.Wait()

	out, err := json.MarshalIndent(index, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("comics.json", out, 0644)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("The file 'comics.json' is successfully created'")
	}

}
