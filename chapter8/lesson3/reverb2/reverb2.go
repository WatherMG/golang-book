/*
Example 8.7
Reverb - это TCP-сервер, имитирующий эхо.
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

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
	input := bufio.NewScanner(conn)
	for input.Scan() {
		go echo(conn, input.Text(), 1*time.Second)
	}
	conn.Close()
}

func echo(c io.Writer, text string, duration time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(text))
	time.Sleep(duration)
	fmt.Fprintln(c, "\t", text)
	time.Sleep(duration)
	fmt.Fprintln(c, "\t", strings.ToLower(text))
}
