/*
Exercise 5.2
Напишите функцию для заполнения map, ключами которого являются имена элементов (р, div, span и т.д.),
а значениями — количество элементов с таким именем в дереве HTML-документа.

Write a function to populate a mapping from element names — p, div, span, and so on — to the number of elements with
that name in an HTML document tree.

Проблемы:
Была проблема с пониманием задачи, сначала решил ее использовав код из ex5.1.go, после понял, что нужно с нуля написать
функцию по другому. Решил задачу с помощью документации к библиотеке https://pkg.go.dev/golang.org/x/net/html
*/

package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func init() {
	tagFreq = make(map[string]int, 0)
}

var tagFreq map[string]int

func main() {
	getTagFreq(os.Stdin)

	for k, v := range tagFreq {
		fmt.Printf("%5d %s\n", v, k)
	}
}

func getTagFreq(r io.Reader) {
	z := html.NewTokenizer(r)

	for {
		typeToken := z.Next()
		tagName, _ := z.TagName()
		if len(tagName) > 0 && typeToken == html.StartTagToken || typeToken == html.SelfClosingTagToken {

			tagFreq[string(tagName)]++
		}
		if typeToken == html.ErrorToken || z.Err() == io.EOF {
			break
		}
	}
}
