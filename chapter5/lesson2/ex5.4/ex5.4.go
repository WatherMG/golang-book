/*
Exercise 5.4

Расширьте функцию visit так, чтобы она извлекала другие разновидности ссылок из документа, такие как изображения,
сценарии и листы стилей.

Extend the visit function so that it extracts other kinds of links from the document, such as images,
scripts, and style sheets.
*/

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex5.4: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit добавляет в links все ссылки, найденные в n, и возвращает результат
func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a", "link":
			links = getLinksFromNode(links, n)
		case "img", "script":
			links = getLinksFromNode(links, n)
		}
	}

	links = visit(links, n.FirstChild)
	return visit(links, n.NextSibling)
}

func getLinksFromNode(l []string, n *html.Node) []string {
	for _, c := range n.Attr {
		if c.Key == "src" || c.Key == "href" {
			l = append(l, c.Key+" "+c.Val)
		}
	}
	return l
}
