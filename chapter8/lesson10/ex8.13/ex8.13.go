/*
Exercise 8.13
Заставьте сервер отключать простаивающих клиентов, которые не
прислали ни одного сообщения за последние 5 минут. Указание: вызов conn.Close()
в другой go-подпрограмме деблокирует активный вызов Read, такой, как
выполняемый вызовом input.Scan().
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

const timeout = 5 * time.Minute

type client struct {
	ch   chan<- string
	name string
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // Все входящие сообщения клиента
)

func broadcaster() {
	clients := make(map[client]bool) // Все подключенные клиенты
	for {
		select {
		case msg := <-messages:
			// Широковещательное входящее сообщение во все
			// каналы исходящих сообщений для клиентов.
			for cli := range clients {
				cli.ch <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			var names []string
			for c := range clients {
				names = append(names, c.name)
			}
			cli.ch <- fmt.Sprintf("%d users in chat: %v", len(clients), names)
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // Исходящие сообщения клиентов
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := client{ch, who}
	ch <- "Вы " + who
	messages <- who + " подключился"
	entering <- cli

	timer := time.NewTimer(timeout)

	go func() {
		<-timer.C
		conn.Close()
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()

	}
	// Примечание: игнорируем потенциальные ошибки input.Err()

	leaving <- cli
	messages <- who + " отключился"

	conn.Close()
}

func clientWriter(conn io.Writer, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintf(conn, msg+"\n") // Примечание: игнорируем ошибки сети
	}
}
