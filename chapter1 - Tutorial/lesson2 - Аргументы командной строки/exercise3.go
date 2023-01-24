/*Поэкспериментируйте с измерением разницы времени выполнения потенциально неэффективных версий и версии с применением strings.Join.*/
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
	fmt.Printf("Time to execute: %s\n", time.Since(start))
}
