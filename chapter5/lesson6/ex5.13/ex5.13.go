/*
Exercise 5.13
Модифицируйте функцию crawl так, чтобы она делала локальные копии найденных ею страниц, при необходимости создавая каталоги.
Не делайте копии страниц, полученных из других доменов. Например, если исходная страница поступает с адреса golang.org,
сохраняйте все страницы оттуда, но не сохраняйте страницы, например, с vimeo.com.

Modify crawl to make local copies of the pages it finds, creating directories as necessary.
Don’t make copies of pages that come from a different domain. For example, if the original page comes from golang.org,
save all files from there, but exclude ones from vimeo.com.
*/

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"GolangBook/chapter5/lesson6/links"
)

var baseHost string

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

func crawl(u string) []string {
	page, err := url.Parse(u)
	if err != nil {
		fmt.Errorf("bad url: %v", err)
		return nil
	}

	if baseHost == "" {
		baseHost = u
	}
	if !strings.Contains(u, baseHost) {
		return nil
	}

	if err := downloadPage(page, u); err != nil {
		fmt.Errorf("can't download page %s: %v", u, err)
	}

	fmt.Println(u)

	list, err := links.Extract(u)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	// Поиск в ширину, начиная с аргумента командной строки
	breadthFirst(crawl, os.Args[1:])
}

func downloadPage(page *url.URL, u string) error {
	dir := page.Host
	var filename string
	if filepath.Ext(page.Path) == "" {
		dir = filepath.Join(dir, page.Path)
		filename = filepath.Join(dir, "index.html")
	} else {
		dir = filepath.Join(dir, filepath.Dir(page.Path))
		filename = filepath.Join(dir, page.Path)
	}

	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return fmt.Errorf("can't create dir %s: %v", dir, err)
	}

	resp, err := http.Get(u)
	if err != nil {
		return fmt.Errorf("can't get %s: %s", resp.Request.URL, resp.Status)
	}

	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("can't create file %s: %v", filename, err)
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("can't write into the file %s: %v", file.Name(), err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("can't close the file %s: %v", file.Name(), err)
		}
	}(file)

	return nil
}
