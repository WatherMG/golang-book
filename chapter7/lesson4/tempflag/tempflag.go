/*
Example 7.4
Tempflag печатает значение своего флага -temp (температура).
*/

package main

import (
	"flag"
	"fmt"

	"GolangBook/chapter7/lesson4/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
