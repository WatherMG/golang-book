/*
Exercise 8.1.
Измените программу clock2 таким образом, чтобы она принимала номер
порта, и напишите программу clockwall, которая действует в качестве клиента
нескольких серверов одновременно, считывая время из каждого и выводя результаты
в виде таблицы, сродни настенным часам, которые можно увидеть в некоторых
офисах. Если у вас есть доступ к географически разнесенным компьютерам,
запустите экземпляры серверов удаленно; в противном случае запустите локальные
экземпляры на разных портах с поддельными часовыми поясами.
*/
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.String("p", "8000", "port of service")
var tz = flag.String("tz", "Europe/Moscow", "Time Zone")

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // Например, обрыв соединения
			continue
		}
		go handleConn(conn, *tz) // Параллельная обработка соединений
		log.Printf("%s is connected", conn.LocalAddr())
	}
}

func handleConn(c net.Conn, loc string) {
	defer c.Close()
	layout, err := time.LoadLocation(loc)
	if err != nil {
		return
	}
	for {
		_, err := io.WriteString(c, time.Now().In(layout).Format("15:04:05\n"))
		if err != nil {
			return // Например, отключение клиента
		}
		time.Sleep(1 * time.Second)
	}
}
