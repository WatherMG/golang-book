package mandelbrot

import (
	"image"
	"image/color"
	"math/cmplx"
	"sync"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func ConcurrentRender(workers int) *image.RGBA {
	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup
	// Создаем изображение с заданными размерами
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	// Создаем канал для передачи номеров строк между горутинами
	size := make(chan int, height)
	// Запускаем анонимную горутину для отправки номеров строк в канал
	go func() {
		for row := 0; row < height; row++ {
			size <- row
		}
		close(size) // Закрываем канал после отправки всех номеров строк
	}()

	// Запускаем заданное количество горутин
	for i := 0; i < workers; i++ {
		wg.Add(1) // Увеличиваем счетчик WaitGroup
		go func() {
			for py := range size { // Получаем номера строк из канала
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					// Точка (px, py) представляет комплексное значение z.
					img.Set(px, py, mandelbrot(z))
				}
			}
			wg.Done() // Уменьшаем счетчик WaitGroup
		}()
	}
	wg.Wait() // Ожидаем завершения всех горутин

	return img
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{Y: 255 - contrast*n}
		}
	}
	return color.Black

}
