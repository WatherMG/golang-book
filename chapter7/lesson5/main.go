package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	fmt.Println(w)
	w = os.Stdout
	fmt.Println(w)
	w = new(bytes.Buffer)
	fmt.Println(w)
	w = nil
	fmt.Println(w)

}
