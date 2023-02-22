/*
Example 4.4
Пример работы со срезом "на лету"
*/
package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty2(data)) // ["one" "three"]
	fmt.Printf("%q\n", data)            // ["one" "three" "three"]
}

// nonempty возвращает срез, содержащий только непустые строки.
// Содержимое базового массива при работе функции изменяется.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] // Срез нулевой длины из исходного среза
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
