/*
Exercise 5.8
Измените функцию forEachNode так, чтобы функции рrе и post возвращали булево значение, указывающее,
следует ли продолжать обход дерева. Воспользуйтесь ими для написания функции ElementBylD с приведенной ниже сигнатурой,
которая находит первый HTML-элемент с указанным атрибутом id. Функция должна прекращать обход дерева,
как только соответствующий элемент найден:

func ElementByID(doc *html.Node, id string) *html.Node


Modify forEachNode so that the pre and post functions return a boolean result indicating whether to continue the traversal.
Use it to write a function ElementByID with the following signature that finds the first HTML element with
the specified id attribute. The function should stop the traversal as soon as a match is found.

func ElementByID(doc *html.Node, id string) *html.Node
*/

package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func init() {
	flag.Parse()
}

var (
	url = flag.String("url", "", "")
	id  = flag.String("id", "", "")
)

// usage go run ex5.8.go -id=site-nav -url https://www.scrapethissite.com/pages/frames/?frame=i
func main() {
	node, err := outline(*url, *id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
		os.Exit(1)
	}

	fmt.Printf("%s ", node.Data)
	for _, a := range node.Attr {
		if a.Key == "id" && a.Val == *id {
			fmt.Printf("%s='%s' found\n", a.Key, a.Val)
		}
	}
}

func outline(url, id string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	node := ElementByID(doc, id)
	if node == nil {
		return nil, errors.New("node not found")
	}
	return node, nil
}

// ElementByID возвращает первый элемент с заданным идентификатором.
// Если элемент не найден, возвращает nil.
func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachElement(doc, id, findElement, findElement)
}

func findElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return true
			}
		}
	}
	return false
}

func forEachElement(n *html.Node, id string, pre, post func(*html.Node, string) bool) *html.Node {
	if pre != nil && pre(n, id) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if node := forEachElement(c, id, pre, post); node != nil {
			return node
		}
	}

	if post != nil && post(n, id) {
		return n
	}

	return nil
}
