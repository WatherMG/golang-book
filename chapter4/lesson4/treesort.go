/*
Example 4.6
Реализация сортировки вставками в бинарном дереве
*/

package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func main() {
	m := []int{9, 3, 4, 1, 90, 344, 123, 7, 55, 7, 978, 0, 2}
	fmt.Println(m)
	Sort(m)
	fmt.Println(m)
}

// Sort сортирует значения "на лету"
func Sort(values []int) {
	var root *tree
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
