/*
Example 4.5
Напишите функцию, которая без выделения дополнительной памяти удаляет все смежные дубликаты в срезе []string.
*/

package main

import "fmt"

func main() {
	s := []string{"a", "a", "a", "b", "b", "c", "c", "a", "a", "b", "b", "c", "c", "a", "a", "b", "b", "c", "c"}
	fmt.Println(unique(s))
}

func unique(strings []string) []string {
	w := 0
	for _, s := range strings {
		if strings[w] == s {
			continue
		}
		w++
		strings[w] = s
	}
	return strings[:w+1]
}
