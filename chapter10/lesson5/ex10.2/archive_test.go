package main

import (
	"log"
	"os"
	"sort"
	"testing"

	"GolangBook/chapter10/lesson5/ex10.2/archive"

	_ "GolangBook/chapter10/lesson5/ex10.2/archive/tar"
	_ "GolangBook/chapter10/lesson5/ex10.2/archive/zip"
)

func TestTar(t *testing.T) {
	// tar
	tf, err := os.Open("cmd/tar.tar")
	if err != nil {
		log.Fatal(err)
	}
	defer tf.Close()
	theaders, err := archive.List(tf)
	if err != nil {
		log.Panic(err)
	}
	var tgot []string
	for _, h := range theaders {
		tgot = append(tgot, h.Name)
	}
	sort.Strings(tgot)

	// zip
	zf, err := os.Open("cmd/zip.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer zf.Close()
	zheaders, err := archive.List(zf)
	if err != nil {
		log.Fatal(err)
	}
	var zgot []string
	for _, h := range zheaders {
		zgot = append(zgot, h.Name)
	}
	sort.Strings(zgot)

	// test
	want := []string{"readme.txt", "gopher.txt", "todo.txt"}
	sort.Strings(want)

	if !equal(tgot, want) {
		t.Errorf("tar reading err")
	}

	if !equal(zgot, want) {
		t.Errorf("zip reading err")
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
