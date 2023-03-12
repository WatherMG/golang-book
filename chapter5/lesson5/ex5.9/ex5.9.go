/*
Exercise 5.9
Напишите функцию expand(s string, f func(string) string) string, которая заменяет каждую подстроку "$foo" в s текстом,
который возвращается вызовом f ("foo").

Write a function expand(s string, f func(string) string) string that replaces each substring “$foo” within s
by the text returned by f('foo').
*/

package main

import (
	"fmt"
	"log"
	"strings"
	"unicode"
)

var text = `Hi there $name.
How is $place? I hope you've been getting a lot of $activity in. Is $someone there? I'm absolutely $expletive going to be there soon.`

func main() {
	fmt.Println(expand(text, replace))
}

func expand(s string, f func(foo string) string) string {
	var result strings.Builder
	for len(s) > 0 {
		i := strings.Index(s, "$")
		if i < 0 {
			result.WriteString(s)
			break
		}
		result.WriteString(s[:i])
		s = s[i+1:]
		j := strings.IndexFunc(s, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_'
		})
		if j < 0 {
			result.WriteString(f(s))
			break
		}
		result.WriteString(f(s[:j]))
		s = s[j:]
	}
	return result.String()
}

func replace(s string) string {
	log.Print(s)
	return strings.ToUpper(s)
}
