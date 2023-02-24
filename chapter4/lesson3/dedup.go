/*
Example 4.3
сохраняет в map только не дублирующие строки
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]bool) // Множество строк
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(seen)
			fmt.Println(line)
		} else {
			fmt.Println(seen)
		}
	}
	if err := input.Err(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
