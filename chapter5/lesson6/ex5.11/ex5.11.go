/*
Exercise 5.11
Преподаватель линейной алгебры (linear algebra) считает, что до его курса следует прослушать курс матанализа (calculus).
Перепишите функцию topoSort так, чтобы она сообщала о наличии циклов.
*/

package main

import (
	"fmt"
	"log"
	"sort"
)

// Uncomment any line below to introduce a dependency cycle.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	// "linear algebra": {"calculus"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	// "computer organization": {"compilers"},
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
	sorted, err := topoSort(prereqs)
	if err != nil {
		log.Println(err)
	}
	for i, course := range sorted {
		fmt.Printf("%2.d: %s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string) error

	visitAll = func(items []string) error {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				if err := visitAll(m[item]); err != nil {
					return err
				}
				order = append(order, item)
			} else {
				hasCycle := true
				for _, s := range order {
					if s == item {
						hasCycle = false
					}
				}
				if hasCycle {
					return fmt.Errorf("has cycle: %s", item)
				}
			}
		}
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	if err := visitAll(keys); err != nil {
		return nil, err
	}

	return order, nil

}
