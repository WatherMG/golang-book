/*
Exercise 4.13
Веб-служба Open Movie Database https://omdbapi.com/ на базе JSON позволяет выполнять поиск фильма по названию и
загружать его афишу. Напишите инструмент poster, который загружает афишу фильма по переданному в командной строке названию.

The JSON-based web service of the Open Movie Database lets you search https://omdbapi.com/ for a movie by name and
download its poster image. Write a tool poster that downloads the poster image for the movie named on the command line.
*/

package movie

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const baseAPI = "https://www.omdbapi.com/"

type Movie struct {
	Poster   string `json:"poster,omitempty"`
	Response string `json:"Response"`
}

func GetFromTitle(t []string) (*Movie, error) {
	q := url.QueryEscape(strings.Join(t, "_"))
	u := fmt.Sprintf("%s?t=%s&apikey=%s", baseAPI, q, os.Getenv("OMDBAPI_KEY"))
	m, err := getMovieFromURL(u)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func GetFromIMDb(t string) (*Movie, error) {
	u := fmt.Sprintf("%s?i=%s&apikey=%s", baseAPI, url.QueryEscape(t), os.Getenv("OMDBAPI_KEY"))
	m, err := getMovieFromURL(u)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func getMovieFromURL(u string) (*Movie, error) {
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("can't get poster for %s, reason: %s", resp.Request.URL, resp.Status)
	}

	defer resp.Body.Close()

	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return nil, fmt.Errorf("can't unmarshaling response: %s", err)
	}
	response, err := strconv.ParseBool(movie.Response)
	if err != nil {
		log.Fatal(err)
	}
	if response {
		return &movie, nil
	} else {
		return nil, fmt.Errorf("movie not found")
	}

}
