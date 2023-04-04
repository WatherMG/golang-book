/*
Exercise 7.17
Расширьте возможности xmlselect так, чтобы элементы могли быть
выбраны не только по имени, но и по атрибутам, в духе CSS, так что, например,
элемент наподобие <div id="page” class="wide"> может быть выбран как по
соответствию атрибутов id или class, так и по его имени.
URL изменился на w3c. актуальный: https://www.w3.org/TR/xml11/.
*/

package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(bytes.NewReader(fetch(os.Args[1])))
	var stack []string            // Стек имен элементов
	var attrs []map[string]string // Срез карт с атрибутами и их значениями. Элемент stack[i] соответствует attrs[i]
	for {
		tok, err := dec.Token()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // Добавление в стек
			attr := make(map[string]string, 0)
			// Добавление в срез атрибута и его значения, даже если их нет.
			// Чтобы атрибуты и значения соответствовали stack[i].
			for _, a := range tok.Attr {
				attr[a.Name.Local] = a.Value
			}
			attrs = append(attrs, attr)
		case xml.EndElement:
			stack = stack[:len(stack)-1] // Удаление из стека
			attrs = attrs[:len(attrs)-1] // Удаление из среза
		case xml.CharData:
			if containsAll(toSlice(stack, attrs), os.Args[2:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// toSlice Добавляет в результирующий срез тег и его атрибуты, для дальнейшего сравнения в containsAll.
func toSlice(stack []string, attrs []map[string]string) (result []string) {
	for i, name := range stack {
		// Добавляем в срез имя тега.
		result = append(result, name)
		for attr, value := range attrs[i] {
			// Добавляем в срез для каждого тега его атрибуты.
			result = append(result, attr+"="+value)
		}

	}
	return result
}

// containsAll указывает, содержит ли x элементы y в том же порядке.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func fetch(uri string) []byte {
	response, err := http.Get(uri)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	body, err := io.ReadAll(response.Body)
	_ = response.Body.Close()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", uri, err)
		os.Exit(1)
	}
	return body
}
