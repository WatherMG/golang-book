/*
Example 8.12
Пакет thumbnail создает изображения размером с миниатюру из более
крупных изображений. В настоящее время поддерживаются только изображения JPEG.
*/

package thumbnail

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Image(src image.Image) image.Image {
	// Вычисление размера миниатюры, сохраняя соотношение сторон.
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y
	width, height := 128, 128
	if aspect := float64(xs) / float64(ys); aspect < 1.0 {
		width = int(128 * aspect) // портретная ориентация
	} else {
		height = int(128 / aspect) // альбомная ориентация
	}

	xscale := float64(xs) / float64(width)
	yscale := float64(ys) / float64(height)

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	// очень грубый алгоритм масштабирования
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcx := int(float64(x) * xscale)
			srcy := int(float64(y) * yscale)
			dst.Set(x, y, src.At(srcx, srcy))
		}
	}
	return dst
}

// ImageStream читает изображение из r и записывает его миниатюру в w
func ImageStream(w io.Writer, r io.Reader) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	dst := Image(src)
	return jpeg.Encode(w, dst, nil)

}

// ImageFile2 считывает изображение из infile и записывает
// его уменьшенную версию в outfile.
func ImageFile2(outfile, infile string) (err error) {
	in, err := os.Open(infile)
	if err != nil {
		return err
	}

	defer in.Close()

	out, err := os.Create(outfile)
	if err != nil {
		return err
	}

	if err := ImageStream(out, in); err != nil {
		out.Close()
		return fmt.Errorf("scaling %s to %s: %w", infile, outfile, err)
	}
	return out.Close()
}

// ImageFile считывает изображение из infile и записывает
// его уменьшенную версию в тот же каталог. Он возвращает
// сгенерированное имя файла, например, "foo.thumb.jpeg".
func ImageFile(infile string) (string, error) {
	ext := filepath.Ext(infile) // e.g. .jpg, .JPEG
	outfile := strings.TrimSuffix(infile, ext) + ".thumb" + ext
	return outfile, ImageFile2(outfile, infile)
}
