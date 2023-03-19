/*
Example 6.2
urlvalues демонстрирует тип карты с методами
*/
package main

import (
	"fmt"
	"net/url"
)

func main() {
	m := url.Values{"lang": {"en"}} // Непосредственное создание
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang")) // en
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // 1 (первое значение)
	fmt.Println(m["item"])     // [1 2] (непосредственное обращение)

	m = nil
	fmt.Println(m.Get("item"))  // ""
	url.Values(nil).Get("item") // "" - эквивалентна предыдущей записи
	m.Add("item", "3")          // panic: присваивание записи в пустой карте.
}
