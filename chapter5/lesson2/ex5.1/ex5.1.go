/*
Exercise 5.1
Измените программу findlinks так, чтобы она обходила связанный список n.FirstChild с помощью рекурсивных вызовов visit,
а не с помощью цикла.

Change the findlinks program to traverse the n.FirstChild linked list using recursive calls to visit instead of a loop.
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
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit добавляет в links все ссылки, найденные в n, и возвращает результат
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

/*
//!+html
package html
type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}
type NodeType int32
const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)
type Attribute struct {
	Key, Val string
}
func Parse(r io.Reader) (*Node, error)
//!-html
*/
