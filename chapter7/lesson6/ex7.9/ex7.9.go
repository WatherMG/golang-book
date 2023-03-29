/*
Exercise 7.9
Воспользуйтесь пакетом html/template (раздел 4.6) для замены printTracks функцией,
которая выводит дорожки в виде таблицы HTML.
Используйте решение предыдущего упражнения для того, чтобы каждый щелчок на заголовке столбца
генерировал HTTP-запрос на сортировку таблицы.
*/

package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

var lastRev bool

func main() {
	http.HandleFunc("/", getRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func printTracks(w http.ResponseWriter, r *http.Request) {
	trackList := template.Must(template.ParseFiles("chapter7/lesson6/ex7.9/index.html"))
	if err := trackList.Execute(w, &tracks); err != nil {
		log.Println(err)
	}
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/title":
		// sort acs or desc on reclick
		if !lastRev {
			sort.Stable(sort.Reverse(sortByColumns(tracks, colTitle)))
			lastRev = true
		} else {
			sort.Stable(sortByColumns(tracks, colTitle))
			lastRev = false
		}
	case "/artist":
		sort.Stable(sortByColumns(tracks, colArtist))
	case "/album":
		sort.Stable(sortByColumns(tracks, colAlbum))
	case "/year":
		sort.Stable(sortByColumns(tracks, colYear))
	case "/length":
		sort.Stable(sortByColumns(tracks, colLength))
	}
	printTracks(w, r)
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
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
