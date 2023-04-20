/*
Exercise 8.8
Используя инструкцию select, добавьте к эхо-серверу из раздела 8.3
тайм-аут, чтобы он отключал любого клиента, который ничего не передает в течение
10 секунд.
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

const timeout = 10 * time.Second

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn io.ReadWriteCloser) {
	wg := &sync.WaitGroup{}
	defer func() {
		wg.Wait()
		conn.Close()
	}()
	input := bufio.NewScanner(conn)

	ch := make(chan string)

	go func() {
		for input.Scan() {
			ch <- input.Text()
		}
	}()

	for {
		select {
		case text := <-ch:
			wg.Add(1)
			go echo(conn, text, 1*time.Second, wg)
		case <-time.After(timeout):
			fmt.Fprintln(conn, "\rYou don't enter message. Exit...")
			return
		}
	}
}

func echo(c io.Writer, text string, duration time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(text))
	time.Sleep(duration)
	fmt.Fprintln(c, "\t", text)
	time.Sleep(duration)
	fmt.Fprintln(c, "\t", strings.ToLower(text))
}
