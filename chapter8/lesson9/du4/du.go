/*
Команда du3 вычисляет суммарный размер всех файлов в каталоге.
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "вывод промежуточных результатов")
var done = make(chan struct{})

func main() {
	// Определяем начальные каталоги.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Параллельный обход дерева файлов.
	fileSizes := make(chan int64)
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, fileSizes)
	}

	// Отмена обхода при обнаруженном вводе.
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	// Ожидание завершения всех горутин и закрытие канала
	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	// Периодический вывод результатов
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	// Вывод информации.
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			// Опустошение канала fileSizes, чтобы позволить
			// завершиться существующим горутинам
			for range fileSizes {
				// Ничего не делаем
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes был закрыт
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskusage(nfiles, nbytes)
		}

	}
	printDiskusage(nfiles, nbytes)
}

func printDiskusage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.3f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir рекурсивно обходит дерево файлов с корнем в dir,
// и отправляет размер каждого найденного файла в fileSizes.
func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()

	if cancelled() {
		return
	}

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, wg, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// sema - подсчитывающий семафор для ограничения параллельности.
var sema = make(chan struct{}, runtime.GOMAXPROCS(-1))

// dirents возвращает записи каталога dir
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // Захват маркера
	case <-done:
		return nil // Отмена
	}

	defer func() {
		<-sema // Освобождение токена
	}()

	// Чтение каталога
	entries, err := os.ReadDir(dir) // ioutils.ReadDir deprecated
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	infos := make([]os.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Fprintf(os.Stderr, "du1: %v\n", err)
			return nil
		}
		infos = append(infos, info)
	}
	return infos
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
