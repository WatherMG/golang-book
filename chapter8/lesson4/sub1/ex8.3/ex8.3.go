/*
Exercise 8.3 В программе netcat3 значение интерфейса conn имеет конкретный тип
*net.TCPConn, который представляет TCP-соединение. ТСР-соединение состоит из
двух половин, которые могут быть закрыты независимо с использованием методов
CloseRead и CloseWrite. Измените главную go-подпрограмму netcat3 так, чтобы она
закрывала только записывающую половину соединения, так, чтобы программа
продолжала выводить последние эхо от сервера reverb1 даже после того, как
стандартный ввод будет закрыт. (Сделать это для сервера reverb2 труднее; см.
упражнение 8.4.)
*/
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
