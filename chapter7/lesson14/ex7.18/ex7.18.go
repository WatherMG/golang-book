/*
Exercise 7.18

С помощью API декодера на основе лексем напишите программу, которая будет читать
произвольный XML-документ и строить представляющее его дерево. Узлы могут быть
двух видов: узлы CharData представляют текстовые строки, а узлы Element —
именованные элементы и их атрибуты. Каждый узел элемента имеет срез дочерних
узлов.

Вам могут пригодиться следующие объявления:

import "encoding/xml"

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type xml.Name
	Attr []xml.Attr
	Children []Node
}
*/

package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	d := bufio.NewReader(fetch(os.Args[1]))
	node, err := parse(d)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", node)
}

type Node interface {
	String() string
}

type CharData string

func (c CharData) String() string {
	return string(c)
}

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (n *Element) String() string {
	b := &bytes.Buffer{}
	visit(n, b, 0)
	return b.String()
}

func visit(n Node, w io.Writer, depth int) {
	var attrs string
	switch n := n.(type) {
	case *Element:
		for _, attr := range n.Attr {
			attrs += fmt.Sprintf(" %s=%q", attr.Name.Local, attr.Value)
		}
		fmt.Fprintf(w, "%*s<%s%s>\n", depth*2, "", n.Type.Local, attrs)
		for _, c := range n.Children {
			visit(c, w, depth+1)
		}
	case CharData:
		fmt.Fprintf(w, "%*s%q\n", depth*2, "", n)
	default:
		panic(fmt.Sprintf("got %T", n))
	}
}

func parse(r io.Reader) (Node, error) {
	dec := xml.NewDecoder(r)
	var stack []*Element
	for {
		t, err := dec.Token()
		if errors.Is(err, io.EOF) {
			return nil, err
		}

		switch t := t.(type) {
		case xml.StartElement:
			el := &Element{t.Name, t.Attr, []Node{}}
			if len(stack) > 0 {
				p := stack[len(stack)-1]
				p.Children = append(p.Children, el)
			}
			stack = append(stack, el)
		case xml.EndElement:
			if len(stack) == 0 {
				return nil, fmt.Errorf("unexpected tag closing")
			} else if len(stack) == 1 {
				return stack[0], nil
			}
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if len(stack) > 0 {
				p := stack[len(stack)-1]
				p.Children = append(p.Children, CharData(t))
			}
		}
	}
}

func fetch(s string) *bytes.Buffer {
	buf := &bytes.Buffer{}
	resp, err := http.Get(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: read %s: %v\n", s, err)
	}
	buf.Write(body)
	return buf
}
