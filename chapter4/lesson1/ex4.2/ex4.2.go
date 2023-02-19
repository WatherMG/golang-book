/*
Exercise 4.2
Напишите программу, которая по умолчанию выводит дайджест SHA256 для входных данных, но при использовании
соответствующих флагов командной строки выводит SHA384 или SHA512.

*/

package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
)

var size = flag.Int("w", 256, "Hash size (256, 384, 512)")
var text = flag.String("t", "asd", "Specify message to get hash")

func main() {
	flag.Parse()
	input := *text

	switch *size {
	case 384:
		fmt.Printf("sha384 for %q: %x\n", input, sha256.Sum256([]byte(input)))
	case 512:
		fmt.Printf("sha512 for %q: %x\n", input, sha256.Sum256([]byte(input)))
	default:
		fmt.Printf("sha256 for %q: %x\n", input, sha256.Sum256([]byte(input)))
	}
}
