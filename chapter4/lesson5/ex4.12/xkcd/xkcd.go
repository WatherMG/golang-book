/*
Exercise 4.11
Популярный веб-ресурс с комиксами xkcd имеет интерфейс JSON. Например, запрос https://xkcd.com/571/info.0.json
возвращает детальное описание комикса 571, одного из многочисленных фаворитов сайта.
Загрузите каждый URL (по одному разу!) и постройте автономный список комиксов. Напишите программу xkcd, которая,
используя этот список, будет выводить URL и описание каждого комикса, соответствующего условию поиска,
заданному в командной строке.

The popular web comic xkcd has a JSON interface. For example, a request to https://xkcd.com/571/info.0.json
produces a detailed description of comic 571, one of many favorites.
Download each URL (once!) and build an offline index. Write a tool xkcd that, using this index,
prints the URL and transcript of each comic that matches a search term provided on the command line
*/

package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const baseAPI = "https://xkcd.com/"

const JSONAPI = "/info.0.json"
const MinNum = 1
const MaxNum = 2743

type Index struct {
	Comics []*Comic
}

type Comic struct {
	Title      string `json:"title,omitempty"`
	SafeTitle  string `json:"safe_title,omitempty"`
	Transcript string `json:"transcript,omitempty"`
	Img        string `json:"img,omitempty"`
	Alt        string `json:"alt,omitempty"`
	News       string `json:"news,omitempty"`
	Link       string `json:"link,omitempty"`
	Year       string `json:"year,omitempty"`
	Month      string `json:"month,omitempty"`
	Day        string `json:"day,omitempty"`
	Num        int    `json:"num,omitempty"`
}

func New(l, c int) Index {
	return Index{make([]*Comic, l, c)}
}

func Get(i int) (*Comic, error) {
	if i == 404 {
		i += 1
	}
	u := fmt.Sprintf("%s%d%s", baseAPI, i, JSONAPI)
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Response error:#%s %s\n", resp.Request.URL, resp.Status)
	}

	defer resp.Body.Close()

	var comic Comic
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return nil, err
	}
	comic.Link = baseAPI + strconv.Itoa(comic.Num)

	return &comic, nil
}

func Search(index Index, keywords []string) []*Comic {
	var res []*Comic
	for _, c := range index.Comics {
		isMatch := true
		for _, term := range keywords {
			if !match(c, term) {
				isMatch = false
			}
			if isMatch {
				res = append(res, c)
			}
		}
	}
	return res
}

func match(c *Comic, term string) bool {
	lowerTerm := strings.ToLower(term)

	return strings.Contains(strings.ToLower(c.Title), lowerTerm) ||
		strings.Contains(strings.ToLower(c.SafeTitle), lowerTerm) ||
		strings.Contains(strings.ToLower(c.Transcript), lowerTerm) ||
		strings.Contains(strings.ToLower(c.Img), lowerTerm) ||
		strings.Contains(strings.ToLower(c.Alt), lowerTerm) ||
		strings.Contains(strings.ToLower(c.News), lowerTerm) ||
		strings.Contains(strings.ToLower(c.Link), lowerTerm) ||
		strings.Contains(c.Year, lowerTerm) ||
		strings.Contains(c.Month, lowerTerm) ||
		strings.Contains(c.Day, lowerTerm)

}

func (c *Comic) String() string {
	return fmt.Sprintf("Comic: #%d\nTitle: %s\nImage: %s\nTranscriprion: %s\nURL: %s\nCreated at: %s-%s-%s",
		c.Num, c.Title, c.Img, c.Transcript, c.Link, c.Day, c.Month, c.Year)

}
