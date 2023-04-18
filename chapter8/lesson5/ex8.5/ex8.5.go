/*
Exercise 8.5
Возьмите существующую последовательную программу, такую как
программа вычисления множества Мандельброта из раздела 3.3 или вычисления
трехмерной поверхности из раздела 3.2, и выполните ее главный цикл параллельно,
с использованием каналов. Насколько быстрее стала работать программа на
многопроцессорной машине? Каково оптимальное количество используемых
go-подпрограмм?
*/

package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"runtime"
	"time"

	"GolangBook/chapter8/lesson5/ex8.5/mandelbrot"
	"GolangBook/chapter8/lesson5/ex8.5/surface"
)

func main() {
	t := time.Now()
	img := mandelbrot.SerialRender()
	f, err := os.Create("./chapter8/lesson5/ex8.5/mandelbrot/mandelbrotSerial.png")
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
	f.Close()
	mst := time.Since(t)
	fmt.Printf("mandelbrot: serial: %d ms\n", mst.Milliseconds())

	t = time.Now()
	workers := runtime.GOMAXPROCS(-1)
	img = mandelbrot.ConcurrentRender(workers)
	f, err = os.Create("./chapter8/lesson5/ex8.5/mandelbrot/mandelbrotConcurrent.png")
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
	f.Close()
	mct := time.Since(t)
	fmt.Printf("mandelbrot: concurrency: %d ms\n", mct.Milliseconds())
	fmt.Printf("diff: %.2fx\n", float64(mst.Milliseconds())/float64(mct.Milliseconds()))

	fmt.Println("------------")

	t = time.Now()
	f, err = os.Create("./chapter8/lesson5/ex8.5/surface/testSerial.svg")
	if err != nil {
		log.Fatal(err)
	}
	data := surface.SerialRender()
	if _, err := f.WriteString(data); err != nil {
		log.Fatal(err)
	}
	f.Close()
	sst := time.Since(t)
	fmt.Printf("surface: serial: %d ms\n", sst.Milliseconds())

	t = time.Now()
	f, err = os.Create("./chapter8/lesson5/ex8.5/surface/testConcurrent.svg")
	if err != nil {
		log.Fatal(err)
	}

	data = surface.ConcurrentRender(runtime.GOMAXPROCS(-1))
	if _, err := f.WriteString(data); err != nil {
		log.Fatal(err)
	}
	f.Close()
	sct := time.Since(t)
	fmt.Printf("surface: concurrency: %d ms\n", sct.Milliseconds())
	fmt.Printf("diff: %.2fx\n", float64(sst.Milliseconds())/float64(sct.Milliseconds()))
}

/*
mandelbrot: serial: 173 ms
mandelbrot: concurrency: 51 ms
diff: 3.39x
------------
surface: serial: 41 ms
surface: concurrency: 6 ms
diff: 6.83x

*/
