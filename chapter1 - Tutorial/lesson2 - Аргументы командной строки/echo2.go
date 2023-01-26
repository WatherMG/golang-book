// Example 1-2-2
// Реализация команды Unix `echo` - выводит в одну строку аргументы, переданные в командной строке
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("Time to execute: %.8fs\n", time.Since(start).Seconds())
}
