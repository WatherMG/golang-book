/*
Example 8.4
Clock2 является TCP-сервером, периодически выводящим время, который может обслуживать несколько клиентов одновременно.
*/
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // Например, обрыв соединения
			continue
		}
		go handleConn(conn) // Параллельная обработка соединений
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // Например, отключение клиента
		}
		time.Sleep(1 * time.Second)
	}
}
