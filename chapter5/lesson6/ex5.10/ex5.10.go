/*
Exercise 5.10
Перепишите topoSort так, чтобы вместо срезов использовались карты, и устраните начальную сортировку.
Убедитесь, что результаты, пусть и недетерминированные, представляют собой корректную топологическую сортировку.

Rewrite topoSort to use maps instead of slices and eliminate the initial sort.
Verify that the results, though nondeterministic, are valid topological orderings.
*/

package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":  {"discrete math"},
	"databases":        {"data structures"},
	"discrete math":    {"intro to programming"},
	"formal languages": {"discrete math"},
	"networks":         {"operating systems"},
	"operating systems": {
		"data structures",
		"computer organization",
	},
	"programming languages": {
		"data structures",
		"computer organization",
	},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%2.d: %s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) map[int]string {
	order := make(map[int]string)
	seen := make(map[string]bool)

	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order[len(order)] = item
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	visitAll(keys)
	return order

}
