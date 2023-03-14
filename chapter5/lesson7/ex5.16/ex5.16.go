/*
Exercise 5.16
Напишите вариативную версию функции strings.Join.
Write a variadic version of strings.Join.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(join("_", "Hello,", "World!"))
	fmt.Println(join("_", "Hello,"))
	fmt.Println(join("_"))

	fmt.Println(strings.Join([]string{"Hello,", "World!"}, "_"))
	fmt.Println(strings.Join([]string{"Hello,"}, "_"))
	fmt.Println(strings.Join([]string{}, "_"))

}

func join(sep string, args ...string) string {
	if len(args) == 0 {
		return ""
	}
	var builder strings.Builder
	for i, s := range args {
		if i > 0 {
			builder.WriteString(sep)
		}
		builder.WriteString(s)
	}
	return builder.String()
}
