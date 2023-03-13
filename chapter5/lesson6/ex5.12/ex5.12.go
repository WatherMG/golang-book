/*
Exercise 5.12
Функции startElement и endElement в chapter5/outline2 совместно используют глобальную переменную depth.
Превратите их в анонимные функции, которые совместно используют локальную переменную функции outline.

The startElement and endElement functions in chapter5/outline2 share a global variable, depth.
Turn them into anonymous functions that share a variable local to the outline function.
*/

package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	_ = outline("https://vk.com")
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	var depth int
	getElement := func(n *html.Node, op bool) {
		if n.Type == html.ElementNode {
			if op {
				fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
				depth++
			} else {
				depth--
				fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
			}
		}
	}

	forEachNode(doc, getElement, getElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node, op bool)) {
	if pre != nil {
		pre(n, true)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n, false)
	}
}
