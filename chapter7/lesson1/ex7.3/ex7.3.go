/*
Exercise 7.3
Напишите метод String для типа *tree из /ch4/treesort.go (раздел 4.4), который показывает последовательность значений в дереве.
*/

package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

var root *tree

func main() {
	m := []int{9, 3, 4, 1, 90, 344, 123, 7, 55, 7, 978, 0, 2}
	Sort(m)
	fmt.Println(root)
}

// Sort сортирует значения "на лету"
func Sort(values []int) {
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)

}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Эквивалентно возврату &tree{value: value}
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	var b bytes.Buffer

	b.WriteRune('[')
	t.appendToBuilder(&b)
	if b.Len() > 0 {
		b.Truncate(b.Len() - 1)
	}
	b.WriteRune(']')

	return b.String()
}

func (t *tree) appendToBuilder(b *bytes.Buffer) {
	if t == nil {
		return
	}
	t.left.appendToBuilder(b)
	b.WriteString(strconv.Itoa(t.value))
	b.WriteRune(' ')
	t.right.appendToBuilder(b)
}
