/*Измените программу echo так, чтобы она выводила индекс и значение каждого аргумента по одному аргументу в строке.*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Println(strconv.Itoa(index) + " " + arg)
	}
}
