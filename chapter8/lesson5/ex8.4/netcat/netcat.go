package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	TCPconn := conn.(*net.TCPConn)

	go func() {
		io.Copy(os.Stdout, conn) // Примечание: игнорируем ошибки
		log.Println("done")
		done <- struct{}{} // Сигнал главной горутине
	}()
	mustCopy(conn, os.Stdin)
	TCPconn.CloseWrite()
	<-done // Ожидание завершения фоновой горутины
	TCPconn.CloseRead()
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
