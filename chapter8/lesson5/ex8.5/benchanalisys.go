/*
Benchanalisys - выполняет анализ результатов бенчмарка, основанный на
меньшем значении ns/op, allocs/op, B/op и большем MB/s.
Выделяет наилучший результат зеленым цветом и пишет во сколько раз повысилась эффективность выполнения кода.
*/
package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type BenchmarkResult struct {
	Line        string
	NsPerOp     float64
	MBPerS      float64
	BPerOp      float64
	AllocsPerOp float64
	Score       float64
}

var bestResult BenchmarkResult
var worstResult BenchmarkResult
var serialResult BenchmarkResult

const reg = `Benchmark.*\s+(\d+)\s+ns/op\s+(\d+\.\d+)\s+MB/s\s+(\d+)\s+B/op\s+(\d+)\s+allocs/op`

func main() {
	pac := []string{"mandelbrot", "surface"}
	// Запускаем бенчмарк с помощью команды go test
	cmd := exec.Command("go", "test", "-bench", "." /*, "-benchtime", "2s"*/)
	cmd.Dir = "./chapter8/lesson5/ex8.5/" + pac[1] // pac[0] - mandelbrot, pac[1] - surface
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running benchmark: %v\n", err)
		os.Exit(1)
	}

	// Регулярное выражение для поиска результатов бенчмарка
	re := regexp.MustCompile(reg)

	// Ищем все совпадения в выводе бенчмарка
	matches := re.FindAllStringSubmatch(out.String(), -1)

	for _, match := range matches {
		line := match[0]
		nsPerOp, _ := strconv.ParseFloat(match[1], 64)
		mbPerS, _ := strconv.ParseFloat(match[2], 64)
		bPerOp, _ := strconv.ParseFloat(match[3], 64)
		allocsPerOp, _ := strconv.ParseFloat(match[4], 64)

		score := nsPerOp*0.4 + allocsPerOp*0.3 + bPerOp*0.2 - mbPerS*0.1

		currentResult := BenchmarkResult{
			Line:        line,
			NsPerOp:     nsPerOp,
			MBPerS:      mbPerS,
			BPerOp:      bPerOp,
			AllocsPerOp: allocsPerOp,
			Score:       score,
		}

		if bestResult.Score == 0 || currentResult.Score < bestResult.Score {
			bestResult = currentResult
		}
		if worstResult.Score == 0 || currentResult.Score > worstResult.Score {
			worstResult = currentResult
		}
		if strings.Contains(currentResult.Line, "Serial") {
			serialResult = currentResult
		}
	}

	speedup := worstResult.Score / bestResult.Score
	serialVSConcurrency := serialResult.Score / bestResult.Score

	printResult(out)
	fmt.Println("Итог, расчет по оценке (score):")
	if speedup == serialVSConcurrency {
		fmt.Printf("Ускорение \033[32mв %.3f раз\033[0m по сравнению с последовательным и лучшим параллельным\n", serialVSConcurrency)
	} else {
		fmt.Printf("Ускорение \033[32mв %.3f раз\033[0m между \u001B[32mлучшим\u001B[0m и \u001B[31mхудшим\u001B[0m\n", speedup)
		fmt.Printf("Ускорение \033[32mв %.3f раз\033[0m по сравнению с последовательным и лучшим параллельным\n", serialVSConcurrency)
	}

}

func printResult(b bytes.Buffer) {
	result := &strings.Builder{}
	for _, line := range strings.Split(b.String(), "\n") {
		if line == worstResult.Line {
			result.WriteString("\033[31m")
			result.WriteString(line)
			result.WriteString("\033[0m\n")
		} else if line == bestResult.Line {
			result.WriteString("\033[32m")
			result.WriteString(line)
			result.WriteString("\033[0m\n")
		} else {
			result.WriteString(line)
			result.WriteByte('\n')
		}
	}
	fmt.Println(result.String())
}

/*
goos: windows
goarch: amd64
pkg: GolangBook/chapter8/lesson5/ex8.5/mandelbrot
cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
BenchmarkSerialRender-16                       8         134524650 ns/op          31.18 MB/s     8388721 B/op    1048578 allocs/op
BenchmarkConcurrentRender/workers=1-16                 8         134573912 ns/op          31.17 MB/s     8398507 B/op    1048583 allocs/op
BenchmarkConcurrentRender/workers=2-16                16          68956656 ns/op          60.83 MB/s     8398910 B/op    1048584 allocs/op
BenchmarkConcurrentRender/workers=4-16                33          36161206 ns/op         115.99 MB/s     8398838 B/op    1048586 allocs/op
BenchmarkConcurrentRender/workers=6-16                48          25341298 ns/op         165.51 MB/s     8399066 B/op    1048589 allocs/op
BenchmarkConcurrentRender/workers=8-16                63          21516546 ns/op         194.93 MB/s     8399193 B/op    1048591 allocs/op
BenchmarkConcurrentRender/workers=10-16               69          19939499 ns/op         210.35 MB/s     8399453 B/op    1048594 allocs/op
BenchmarkConcurrentRender/workers=12-16               66          18743564 ns/op         223.77 MB/s     8399359 B/op    1048596 allocs/op
BenchmarkConcurrentRender/workers=16-16               74          16647647 ns/op         251.95 MB/s     8400150 B/op    1048601 allocs/op
BenchmarkConcurrentRender/workers=16#01-16            74          16741630 ns/op         250.53 MB/s     8399173 B/op    1048599 allocs/op
BenchmarkConcurrentRender/workers=24-16               66          16711138 ns/op         250.99 MB/s     8399451 B/op    1048607 allocs/op
BenchmarkConcurrentRender/workers=32-16               76          16663587 ns/op         251.70 MB/s     8399746 B/op    1048616 allocs/op
BenchmarkConcurrentRender/workers=64-16               74          16564135 ns/op         253.22 MB/s     8401098 B/op    1048648 allocs/op
BenchmarkConcurrentRender/workers=128-16              74          17080043 ns/op         245.57 MB/s     8403668 B/op    1048716 allocs/op
BenchmarkConcurrentRender/workers=256-16              69          17041228 ns/op         246.13 MB/s     8407790 B/op    1048846 allocs/op
PASS
ok      GolangBook/chapter8/lesson5/ex8.5/mandelbrot    18.921s

Сравнение скорости последовательного выполнения и наилучшего параллельного:
Ускорение в 8.12 раз
*/
