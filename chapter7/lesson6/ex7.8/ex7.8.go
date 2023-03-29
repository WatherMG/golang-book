/*
Exercise 7.8
Многие графические интерфейсы предоставляют таблицы с многоуровневой сортировкой с сохранением состояния:
первичный ключ определяется по последнему щелчку на заголовке, вторичный — по-предпоследнему и т.д.
Определите реализацию sort.Interface для использования в такой таблице.
Сравните этот подход с многократной сортировкой с использованием sort.Stable.


При использовании sort.Sort() элементы сортируются по одному значению. Это означает, что если два элемента имеют
одинаковое значение, то порядок их следования в отсортированном списке может быть произвольным.

При использовании sort.Stable() элементы сортируются по нескольким значениям. Если два элемента имеют одинаковые
значения для всех полей, по которым они сортируются, то порядок их следования в отсортированном списке
будет таким же, как и в исходном списке.
*/

package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func tracks() []*Track {
	return []*Track{
		{"Go", "Delilah", "From the Roots Up1", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("10m37s")},
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		{"Go", "Delilah1", "From the Roots Up", 2012, length("3m38s")},
	}
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")

	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

func colTitle(x, y *Track) bool  { return x.Title < y.Title }
func colArtist(x, y *Track) bool { return x.Artist < y.Artist }
func colAlbum(x, y *Track) bool  { return x.Album < y.Album }
func colYear(x, y *Track) bool   { return x.Year < y.Year }
func colLength(x, y *Track) bool { return x.Length < y.Length }

type less func(x, y *Track) bool

func sortByColumns(t []*Track, f ...less) *customSort {
	return &customSort{
		tracks:  t,
		columns: f,
	}
}

type customSort struct {
	tracks  []*Track
	columns []less
}

func (x *customSort) Len() int { return len(x.tracks) }
func (x *customSort) Less(i, j int) bool {
	p, q := x.tracks[i], x.tracks[j]
	var k int
	// сравниваем столбцы один за другим, кроме последнего
	for k = 0; k < len(x.columns)-1; k++ {
		f := x.columns[k]
		switch {
		case f(p, q):
			return true
		case f(q, p):
			return false
		}
	}
	// все столбцы равны, используем последнюю колонку как окончательное результат
	return x.columns[k](x.tracks[i], x.tracks[j])
}

func (x *customSort) Swap(i, j int) { x.tracks[i], x.tracks[j] = x.tracks[j], x.tracks[i] }

func useSortByColumns() []*Track {
	t := tracks()
	sort.Sort(sortByColumns(t, colTitle, colArtist))
	return t
}

func useSortStable() []*Track {
	t := tracks()
	sort.Stable(sortByColumns(t, colTitle, colArtist))
	return t
}

func main() {
	fmt.Println("By Title, Artist, Year")
	printTracks(useSortByColumns())

	fmt.Println("\nUse sort.Stable. By Title, Artist, Year")
	printTracks(useSortStable())
}
