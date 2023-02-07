/*
Exercise 1-2-2
Измените программу echo так, чтобы она выводила индекс и значение каждого аргумента по одному аргументу в строке.
*/
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for index, arg := range os.Args[1:] {
		fmt.Printf("%d: %s\n", index+1, arg)
	}
	fmt.Printf("Time to execute: %.8fs\n", time.Since(start).Seconds())
}
