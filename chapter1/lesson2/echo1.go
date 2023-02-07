// Example 1-2-1
// Реализация команды Unix `echo` - выводит в одну строку аргументы, переданные в командной строке
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("Time to execute: %.8fs\n", time.Since(start).Seconds())
}
