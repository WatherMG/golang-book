/*
Example 1-7-2

Server2 - минимальный "echo"-сервер со счетчиком запросов.
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler2) // Каждый запрос вызывает обработчик
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Обработчик, возвращающий компонент пути запрашиваемого URL
func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	_, _ = fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// Счетчик, возвращающий количество сделанных запросов
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	_, _ = fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
