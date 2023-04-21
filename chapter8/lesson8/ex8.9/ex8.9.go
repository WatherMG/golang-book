/*
Напишите версию du, которая вычисляет и периодически выводит отдельные итоговые величины для каждого из каталогов root.
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

type dir struct {
	id   int
	size int64
}

func main() {
	// Определяем начальные каталоги.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// Параллельный обход дерева файлов.
	info := make(chan dir)
	var wg sync.WaitGroup
	for id, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, id, info)
	}
	go func() {
		wg.Wait()
		close(info)
	}()
	// Периодический вывод результатов
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	// Вывод информации.
	nfiles := make([]int64, len(roots))
	nbytes := make([]int64, len(roots))
loop:
	for {
		select {
		case d, ok := <-info:
			if !ok {
				break loop // info был закрыт
			}
			nfiles[d.id]++
			nbytes[d.id] += d.size
		case <-tick:
			printDiskusage(roots, nfiles, nbytes)
		}
	}
	printDiskusage(roots, nfiles, nbytes)
}

func printDiskusage(roots []string, nfiles, nbytes []int64) {
	for id, root := range roots {
		fmt.Printf("%d files, %.1f GB в %s\n", nfiles[id], float64(nbytes[id])/1e9, root)
	}
}

// walkDir рекурсивно обходит дерево файлов с корнем в dir,
// и отправляет размер каждого найденного файла в fileSizes.
func walkDir(d string, wg *sync.WaitGroup, root int, info chan<- dir) {
	defer wg.Done()
	for _, entry := range dirents(d) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(d, entry.Name())
			go walkDir(subdir, wg, root, info)
		} else {
			info <- dir{root, entry.Size()}
		}
	}
}

// sema - подсчитывающий семафор для ограничения параллельности.
var sema = make(chan struct{}, runtime.GOMAXPROCS(-1))

// dirents возвращает записи каталога dir
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{} // Захват токена
	defer func() {
		<-sema // Освобождение токена
	}()
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
