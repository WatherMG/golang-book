// Example 1-3-3
// Реализация команды Unix `echo` - выводит в одну строку аргументы, переданные в командной строке

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("Time to execute: %.8fs\n", time.Since(start).Seconds())
}
