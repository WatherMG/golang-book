/*
Exercise 1-2-3
Поэкспериментируйте с измерением разницы времени выполнения потенциально неэффективных версий и версии с
применением strings.Join.
*/
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	for index, arg := range os.Args[1:] {
		fmt.Println(strconv.Itoa(index) + " " + arg)
	}
	fmt.Printf("Time to execute: %.8fs\n", time.Since(start).Seconds())
}
