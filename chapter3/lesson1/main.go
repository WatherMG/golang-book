package main

import "fmt"

func main() {
	for i := 0; i < 128; i++ {
		fmt.Printf("%d = %08b\n", i<<1, i<<1)
	}
}
