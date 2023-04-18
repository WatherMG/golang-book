package surface

import (
	"fmt"
	"strings"
	"sync"
)

type cell struct {
	i, j int
}
type result struct {
	order int
	data  string
}

var wg sync.WaitGroup

func ConcurrentRender(workers int) string {
	s := make([]string, cells*cells)
	cellsChan := make(chan cell, cells*cells)
	resultChan := make(chan result)

	go func() {
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				cellsChan <- cell{i, j}
			}
		}
		close(cellsChan)
	}()

	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := range cellsChan {
				ax, ay := corner(c.i+1, c.j)
				bx, by := corner(c.i, c.j)
				cx, cy := corner(c.i, c.j+1)
				dx, dy := corner(c.i+1, c.j+1)
				data := fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
				resultChan <- result{order: c.i*cells + c.j, data: data}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for res := range resultChan {
		s[res.order] = res.data
	}

	return fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg'> "+
		"style='stroke: gray; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>%s</svg>", width, height, strings.Join(s, ""))
}
