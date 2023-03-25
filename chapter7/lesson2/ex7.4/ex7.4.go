/*
Exercise 7.4
Функция strings.NewReader возвращает значение, соответствующее интерфейсу io.Reader (и другим),
путем чтения из своего аргумента, который представляет собой строку. Реализуйте простую версию NewReader и используйте
ее для создания синтаксического анализатора HTML (раздел 5.2), принимающего входные данные из строки.
*/

package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type URLReader struct {
	url string
	i   int64
}

func (u *URLReader) Read(b []byte) (n int, err error) {
	if u.i >= int64(len(u.url)) {
		return 0, io.EOF
	}
	n = copy(b, u.url[u.i:])
	u.i = int64(n)
	// fmt.Printf("Parsed url: %s\n", u.url)
	return
}

func NewReader(s string) io.Reader {
	return &URLReader{s, 0}
}

func main() {
	doc, err := html.Parse(NewReader("vk.com"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // Внесение дескриптора в стек
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
