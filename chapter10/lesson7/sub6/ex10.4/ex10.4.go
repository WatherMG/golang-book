/*
Создайте инструмент, который сообщает о множестве всех пакетов в рабочей
области, которые транзитивно зависят от пакетов, указанных аргументами командной
строки. Указание: вам нужно будет выполнить go list дважды: один раз — для
исходных пакетов и один раз — для всех пакетов. Вы можете проанализировать вывод
в формате JSON с помощью пакета encoding/json (раздел 4.5).

*/

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os/exec"
)

type pkg struct {
	ImportPath string
	Deps       []string
}

func getPackages() ([]pkg, error) {
	out, err := exec.Command("go", "list", "-json", "all").Output()
	if err != nil {
		return nil, err
	}
	var packages []pkg
	dec := json.NewDecoder(bytes.NewReader(out))
	for {
		var pkg pkg
		if err := dec.Decode(&pkg); errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		packages = append(packages, pkg)
	}
	return packages, nil
}

func findDependentPackages(packages []pkg, targetDeps []string) []string {
	var result []string
	targets := make(map[string]bool)

	for _, target := range targetDeps {
		targets[target] = true
	}

	for _, pkg := range packages {
		if containsPackage(targets, pkg.Deps) {
			result = append(result, pkg.ImportPath)
		}
	}
	return result
}

func containsPackage(targets map[string]bool, deps []string) bool {
	for _, dep := range deps {
		if targets[dep] {
			return true
		}
	}
	return false
}

func main() {
	flag.Parse()
	targetDeps := flag.Args()

	packages, err := getPackages()
	if err != nil {
		fmt.Println(err)
		return
	}

	results := findDependentPackages(packages, targetDeps)

	for _, result := range results {
		fmt.Println(result)
	}
}
