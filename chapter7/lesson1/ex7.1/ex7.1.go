/*
Exercise 7.1
Используя идеи из ByteCounter, реализуйте счетчики для слов и строк. Вам пригодится функция bufio.ScanWords.
*/

package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type ByteCounter int
type WordCounter int
type LineCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (w *WordCounter) Write(p []byte) (int, error) {
	count := getCount(p, bufio.ScanWords)
	*w += WordCounter(count)
	return len(p), nil
}

func (l *LineCounter) Write(p []byte) (int, error) {
	count := getCount(p, bufio.ScanLines)
	*l += LineCounter(count)
	return len(p), nil
}

func getCount(p []byte, fn bufio.SplitFunc) (result int) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(fn)
	for scanner.Scan() {
		result++
	}
	return result
}

func main() {
	var b ByteCounter
	var w WordCounter
	var l LineCounter

	input := []byte("one\n two\n three\n four\n five")
	b.Write(input)
	w.Write(input)
	l.Write(input)

	// Multiline input from os.Stdin
	/*input := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the multiline text: > ")
	for input.Scan() {
		if input.Text() == "" {
			break
		}
		b.Write([]byte(input.Text()))
		w.Write([]byte(input.Text()))
		l.Write([]byte(input.Text()))

		if err := input.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "reading error: %v", err)
			os.Exit(1)
		}
	}*/
	fmt.Printf("The text have: %d chars, %d words, %d lines\n", b, w, l)
}
