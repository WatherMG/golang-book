/*
Exercise 1-2-1
Измените программу echo так, чтобы она выводила также os. Args[0], имя выполняемой команды.
*/
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[0:], " "))
	fmt.Printf("Time to execute: %.8fs\n", time.Since(start).Seconds())
}
