/*
Exercise 5.7
Разработайте startElement и endElement для обобщенного вывода HTML.
Выводите узлы комментариев, текстовые узлы и атрибуты каждого элемента (<а href='...'>).
Используйте сокращенный вывод наподобие <img/> вместо <img/img>, когда элемент не имеет дочерних узлов.
Напишите тестовую программу, чтобы убедиться в корректности выполняемого анализа (см. главу 11, “Тестирование”.)

Develop startElement and endElement into a general HTML pretty-printer.
Print comment nodes, text nodes, and the attributes of each element (<a href='...'>).
Use short forms like <img/> instead of <img></img> when an element has no children.
Write a test to ensure that the output can be parsed successfully.

Проблемы:
Была проблема с тестированием, а конкретно с получением форматированного вывода для дальнейшей проверки на корректность
выполняемого анализа. Решил эту проблему использованием *bytes.Buffer и fmt.Fprintf(buf, ""), которые: создает буфер,
записывает в буфер форматированную строку, соответственно. Проверка в тесте выполняется с помощью buf.String(), которая
преобразует буфер в строку. Из-за решения этой проблемы пришлось передавать в каждую функцию буфер. Должен быть способ
проще (интерфейсы???). Возможно перепишу программу после прохождения главы 7.
*/

package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

var depth int

func main() {
	if err := prettyPrint("https://learngitbranching.js.org/?locale=ru_RU&NODEMO="); err != nil {
		log.Fatal(err)
	}
}

func prettyPrint(u string) error {
	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	forEachNode(doc, startElement, endElement, buf)
	fmt.Printf("%s", buf.String())
	return nil
}

// forEachNode вызывает функции pre(x) и post(x) для каждого узла х в дереве с корнем n.
func forEachNode(n *html.Node, pre, post func(n *html.Node, buf *bytes.Buffer), buf *bytes.Buffer) {
	if pre != nil {
		pre(n, buf)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post, buf)
	}
	if post != nil {
		post(n, buf)
	}
}

func startElement(n *html.Node, buf *bytes.Buffer) {
	switch n.Type {
	case html.ElementNode:
		getElement(n, buf)
	case html.TextNode:
		getText(n, buf)
	case html.CommentNode:
		getComment(n, buf)
	}
}

func endElement(n *html.Node, buf *bytes.Buffer) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Fprintf(buf, "%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

func getElement(n *html.Node, buf *bytes.Buffer) {
	end := ">"
	if n.FirstChild == nil {
		end = "/>"
	}
	attrs := ""
	if n.Attr != nil {
		for _, a := range n.Attr {
			attrs += fmt.Sprintf(" %s=%q", a.Key, a.Val)
		}
	}
	fmt.Fprintf(buf, "%*s<%s%s%s\n", depth*2, "", n.Data, attrs, end)
	depth++
}

func getText(n *html.Node, buf *bytes.Buffer) {
	text := strings.TrimSpace(n.Data)
	if len(text) != 0 {
		fmt.Fprintf(buf, "%*s%s\n", depth*2, "", text)
	}
}

func getComment(n *html.Node, buf *bytes.Buffer) {
	fmt.Fprintf(buf, "%*s<!--%s-->\n", depth*2, "", n.Data)
}
