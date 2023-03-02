/*
Exercise 4.13
Веб-служба Open Movie Database https://omdbapi.com/ на базе JSON позволяет выполнять поиск фильма по названию и
загружать его афишу. Напишите инструмент poster, который загружает афишу фильма по переданному в командной строке названию.

The JSON-based web service of the Open Movie Database lets you search https://omdbapi.com/ for a movie by name and
download its poster image. Write a tool poster that downloads the poster image for the movie named on the command line.
*/

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"GolangBook/chapter4/lesson5/ex4.13/poster"
)

func init() {
	flag.Parse()
}

var (
	t      []string
	i      = flag.String("i", "", "Movie's IMDb number e.g 'tt1285016'")
	apiKey = flag.String("key", "", "Api key")
)

// usage: run main.go -t MOVIE_NAME || -i MOVIE_IMDb_NUMBER && -key=1239871 if key is not in ENV
func main() {
	t = flag.Args()
	if os.Getenv("OMDBAPI_KEY") == "" {
		err := os.Setenv("OMDBAPI_KEY", *apiKey)
		if err != nil {
			log.Fatal(err)
		}
	}
	if len(flag.Args()) != 0 {

		m, err := poster.GetFromTitle(t)
		if err != nil {
			log.Fatal(err)
		}
		downloadPoster(m)
	} else if len(*i) != 0 {
		m, err := poster.GetFromIMDb(*i)
		if err != nil {
			log.Fatal(err)
		}
		downloadPoster(m)
	} else {
		fmt.Printf("You don't specify movie name or IMDb number")
		os.Exit(1)
	}
}

func downloadPoster(m *poster.Movie) {
	re := regexp.MustCompile("[^a-zA-Z0-9-_]+")
	q := re.ReplaceAllString(strings.Join(t, " "), " ")

	resp, err := http.Get(m.Poster)
	if err != nil {
		fmt.Printf("the poster for the movie '%s' not found in OMDBAPI.com, %s", q, err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	dir := "chapter4/lesson5/ex4.13/images/" + q
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatal(err)
	}

	filename := filepath.Base(q + filepath.Ext(m.Poster))
	path := filepath.Join(dir, filename)
	out, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err == nil {
		fmt.Printf("File %s is created in %s\n", filename, path)
	} else {
		log.Fatal(err)
	}
}
