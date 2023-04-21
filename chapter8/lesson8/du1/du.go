/*
Команда du1 вычисляет суммарный размер всех файлов в каталоге.
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Определяем начальные каталоги.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// Обход дерева файлов.
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// Вывод информации.
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskusage(nfiles, nbytes)
}

func printDiskusage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.3f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir рекурсивно обходит дерево файлов с корнем в dir,
// и отправляет размер каждого найденного файла в fileSizes.
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}

}

// dirents возвращает записи каталога dir
func dirents(dir string) []os.FileInfo {
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
