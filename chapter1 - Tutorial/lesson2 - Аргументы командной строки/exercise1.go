/*Измените программу echo так, чтобы она выводила также os. Args[0], имя выполняемой команды.*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
