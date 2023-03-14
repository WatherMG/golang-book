/*
Exercise 5.17
Напишите вариативную функцию ElementsByTagName, которая для данного дерева узла HTML и нуля или нескольких имен
возвращает все элементы, которые соответствуют одному из этих имен.

Write a variadic function ElementsByTagName that, given an HTML node tree and zero or more names, returns all the
elements that match one of those names.
Here are two example calls:

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node

images := ElementsByTagName(doc, 'img')
headings := ElementsByTagName(doc, 'h1', 'h2', 'h3', 'h4')
*/

package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://www.scrapethissite.com/pages/frames/?frame=i")
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Println(err)
	}

	nodes := ElementsByTagName(doc, "img", "h3", "a")
	for _, n := range nodes {
		fmt.Println(n.Data)
		fmt.Println(n.Attr)
	}
}

func ElementsByTagName(doc *html.Node, tags ...string) (nodes []*html.Node) {
	if len(tags) == 0 {
		return nil
	}
	if doc.Type == html.ElementNode {
		for _, tag := range tags {
			if doc.Data == tag {
				nodes = append(nodes, doc)
			}
		}
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, ElementsByTagName(c, tags...)...)
	}
	return nodes
}
