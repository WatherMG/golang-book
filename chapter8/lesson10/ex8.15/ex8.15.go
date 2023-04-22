/*
Exercise 8.15
Ошибка своевременного чтения данных любой клиентской программой в
конечном итоге вызывает сбой всех клиентских программ. Измените широковещатель
таким образом, чтобы в случае, когда горутина передачи сообщения клиентам
не готова принять сообщение, он вместо ожидания пропускал сообщение. Кроме того,
добавьте буферизацию каждого канала исходящих сообщений клиента, чтобы
большинство сообщений не удалялись. Широковещатель должен использовать
неблокирующее отправление в этот канал.
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
				select {
				case cli.ch <- msg: // Передача сообщения
				case <-time.After(10 * time.Second): // пропуск сообщения
				default: // пропуск сообщения
				}
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

	fmt.Fprintf(conn, "Your name: ")

	sc := bufio.NewScanner(conn)
	sc.Scan()
	who := sc.Text()

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
