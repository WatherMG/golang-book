/*
Exercise 5.3
Напишите функцию для вывода содержимого всех текстовых узлов в дереве документа HTML. Не входите в элементы
<script> и <style>, поскольку их содержимое в веб-браузере не является видимым.

Write a function to print the contents of all text nodes in an HTML document tree. Do not descend into
<script> or <style> elements, since their contents are not visible in a web browser.

Проблемы:
Сначала решил написать функцию с без использования рекурсии, с использованием NewTokenizer.
Столкнулся с проблемой лишних проверок на принадлежность элемента к script и style и удаления их из среза,
чтобы был последовательный вывод: <tag>Text</tag>. Решил попробовать подход с использованием Node,
но пришлось использовать рекурсию.
*/

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	for _, text := range getTextFromHTML(nil, doc) {
		fmt.Println(text)
	}

}

func getTextFromHTML(texts []string, n *html.Node) []string {
	if n == nil {
		return texts
	}
	if n.Type == html.TextNode && n.Parent.Data != "script" && n.Parent.Data != "style" {
		if len(strings.TrimSpace(n.Data)) != 0 {
			for _, line := range strings.Split(n.Data, "\n") {
				if len(line) != 0 {
					texts = append(texts, strings.TrimSpace(line))
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		texts = getTextFromHTML(texts, c)
	}
	return texts
}
