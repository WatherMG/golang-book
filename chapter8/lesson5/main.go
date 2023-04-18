/*
Команда thumbnail создает эскизы файлов JPEG. имена которых указываются в каждой
строке стандартного ввода. Метка "+build ignore" (см. стр.295) исключает этот
файл из пакета thumbnail, но он может быть скомпилирован как команда и запущен
как вот так:
Run with:
$ go run $GOPATH/src/.../thumbnail/main.go
foo.jpeg
^D
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"GolangBook/chapter8/lesson5/thumbnail"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		thumb, err := thumbnail.ImageFile(input.Text())
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}
