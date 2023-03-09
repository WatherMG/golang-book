/*
Exercise 5.5
Реализуйте функцию countWordsAndImages (см. разделение на слова в упр. 4.9).

Узнал разницу работы strings.Split и strings.Fields. Первый принимает в качестве аргумента любую строку как разделитель,
второй использует unicode.IsSpace и unicode.IsControl для определения пробелов и управляющих конструкций. Оба возвращают
срез подстрок из строки.

Проблемы:
Столкнулся с проблемой передачи счетчиков в вызывающую функцию из-за того, что при рекурсии счетчик сбрасывался.
Решил с помощью указателей, но показалось, что такой код слишком сложен для чтения. Переписал без использования указателей.
*/

package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	w, i, _ := CountWordsAndImages("https://golang.org")
	fmt.Printf("%4.d words\n%4.d images", w, i)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	defer resp.Body.Close()

	words, images = countWordsAndImages(doc)
	return words, images, nil
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}

	switch n.Type {
	case html.ElementNode:
		if n.Data == "img" {
			images++
		}
	case html.TextNode:
		if len(strings.TrimSpace(n.Data)) > 0 && n.Parent.Data != "script" && n.Parent.Data != "style" {
			words += len(strings.Fields(n.Data))
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}

	return words, images
}

// Вариант работы кода с указателями
/*func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	defer resp.Body.Close()

	words, images = countWordsAndImages(doc, &words, &images) // !
	return words, images, nil
}

func countWordsAndImages(n *html.Node, w, i *int) (words, images int) { // !
	if n == nil {
		return
	}

	switch n.Type {
	case html.ElementNode:
		if n.Data == "img" {
			*i++ // !
		}
	case html.TextNode:
		if len(strings.TrimSpace(n.Data)) > 0 && n.Parent.Data != "script" && n.Parent.Data != "style" {
			*w += len(strings.Fields(n.Data)) // !
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countWordsAndImages(c, w, i) // !
	}

	return *w, *i // !
}*/
