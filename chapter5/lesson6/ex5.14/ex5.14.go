/*
Exercise 5.14
Используйте функцию breadthFirst для исследования другой структуры. Например, вы можете использовать зависимости
учебных курсов из примера topoSort (ориентированный граф), иерархию файловой системы на своем компьютере (дерево) или
список маршрутов автобусов в своем городе (неориентированный граф).
*/

package main

import (
	"fmt"
	"math/rand"
	"sort"
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

// breadthFirst вызывает f для каждого элемента в worklist.
// Все элементы, возвращаемые f, добавляются в worklist.
// f вызывается для каждого элемента не более одного раза
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}

}

// Gets a random key (course) from prereq and outputs the course dependencies, including their dependencies
func main() {
	keys := make([]string, len(prereqs))
	i := 0
	for k := range prereqs {
		keys[i] = k
		i++
	}
	randKey := keys[rand.Intn(len(keys))]

	fmt.Printf("For '%s' you need learn this courses:\n", randKey)
	breadthFirst(getDeps, prereqs[randKey])

}

func getDeps(item string) (keys []string) {
	fmt.Println(item)
	keys = append(keys, prereqs[item]...)
	sort.Strings(keys)
	return keys
}
