/*
Расширьте программу jpeg так, чтобы она преобразовывала любой поддерживаемый
входной формат в любой выходной с использованием функции image.Decode для
определения входного формата и флага командной строки для выбора выходного
формата.
*/

package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"
)

var format = flag.String("f", "jpeg", "output format, required {png, jpeg, gif}")

func main() {
	flag.Parse()
	if len(flag.Args()) > 0 {
		fmt.Fprintln(os.Stderr, "usage: ./ex10.1 -f=png|jpeg|gif <INPUT >OUTPUT")
		os.Exit(1)
	}
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Входной формат =", kind)
	switch strings.ToLower(*format) {
	case "jpeg", "jpg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, &gif.Options{})
	default:
		return fmt.Errorf("unknown output format:%q", *format)
	}

}
