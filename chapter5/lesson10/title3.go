/*
Example 5.18

// Title3 печатает заголовок HTML-документа, указанного по URL.
*/

package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, arg := range os.Args[1:] {
		if err := title(arg); err != nil {
			fmt.Fprintf(os.Stderr, "title: %v\n", err)
		}
	}
}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check Content-Type is HTML (e.g., "text/html; charset=utf8").
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	title, err := soleTitle(doc)
	if err != nil {
		return err
	}
	fmt.Println(title)
	return nil
}

// soleTitle возвращает ошибку первого непустого элемента title
// в документе и ошибку, если их количество не равно одному.
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			// Ожидаемая паника
			err = fmt.Errorf("несколько элементов title")
		default:
			panic(p) // Неожиданная аварийная ситуация; восстанавливаем ее
		}
	}()
	// Прекращаем рекурсию при нескольких непустых элементах title.
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // несколько элементов title
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("нет элемента title")
	}
	return title, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
