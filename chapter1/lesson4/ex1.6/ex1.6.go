/*/*
Exercise 1.6

Генерирует анимированный GIF из случайных фигур Лиссажу.
*/

package main

import (
	"bytes"
	"image"
	color "image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

var colorGreen = color.RGBA{G: 222, A: 255}
var colorBlue = color.RGBA{B: 255, A: 255}
var colorRed = color.RGBA{R: 255, A: 255}

var palette = []color.Color{color.Black, colorGreen, colorBlue, colorRed} // Инстанцирование среза

const (
	blackIndex = 0
	greenIndex = 1
	blueIndex  = 2
	redIndex   = 3
)

func setColorIndex(frameNum int) uint8 {
	colorIndex := uint8(blackIndex)
	switch {
	case frameNum%9 <= 2:
		colorIndex = greenIndex
	case frameNum%9 >= 3 && frameNum%9 <= 5:
		colorIndex = redIndex
	case frameNum%9 >= 6 && frameNum%9 <= 8:
		colorIndex = blueIndex
	}
	return colorIndex
}

func main() {
	buf := &bytes.Buffer{}
	lissajous(buf)
	if err := os.WriteFile(os.Args[1], buf.Bytes(), os.ModePerm); err != nil {
		panic(err)
	}
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Количество полных колебаний по x
		res     = 0.001 // Угловое разрешение
		size    = 100   // Размер канвы изображения [-size..+size]
		nframes = 64    // Количество фреймов
		delay   = 8     // Задержка между кадрами (в *10мс)
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0        // Относительная частота колебаний у
	anim := gif.GIF{LoopCount: nframes} // Инстанцирование структуры
	phase := 0.0                        // Разность фаз
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), setColorIndex(i))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		log.Fatal(err)
	}
}
