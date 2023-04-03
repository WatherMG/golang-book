/*
Example 7.12
Xmlselect выводит текст из выбранных элементов XML-документа.

URL изменился на w3c. актуальный: https://www.w3.org/TR/xml11/
*/

package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string // Стек имен элементов
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
		case xml.EndElement:
			stack = stack[:len(stack)-1] // Удаление из стека
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
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
