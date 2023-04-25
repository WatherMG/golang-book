/*
Пакет memotest предоставляет общие функции для
тестирования различных конструкций пакета memo.
*/

package memotest

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func httpGetBody(url string, done <-chan struct{}) (interface{}, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	go func() {
		select {
		case <-done:
			fmt.Printf("httpGetBody: cancel task %s\n", req.URL.Path)
			cancel()
		case <-ctx.Done():
			fmt.Println("httpGetBody: Done request!")
		}
	}()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

var HTTPGetBody = httpGetBody

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string, done <-chan struct{}) (interface{}, error)
}

func Sequential(t *testing.T, m M, done <-chan struct{}) {
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url, done)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

func Concurrent(t *testing.T, m M, done <-chan struct{}) {
	var wg sync.WaitGroup
	for url := range incomingURLs() {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			start := time.Now()
			value, err := m.Get(u, done)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n", u, time.Since(start), len(value.([]byte)))
		}(url)
	}
	wg.Wait()
}

func ConcurrentCancel(t *testing.T, m M, done chan struct{}) {
	var n sync.WaitGroup

	timer := time.NewTimer(400 * time.Millisecond) // 1.6s
	go func() {
		<-timer.C
		close(done)
		timer.Stop()
	}()

	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url, done)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}
