/*
Exercise 8.2
Реализуйте параллельный FTP-сервер. Сервер должен интерпретировать
команды от каждого клиента, такие как cd для изменения каталога, ls для вывода
списка файлов в каталоге, get для отправки содержимого файла и close для
закрытия соединения. В качестве клиента можно использовать стандартную команду
ftp или написать собственную программу.
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go mustCopy(os.Stdout, conn)

	sc := bufio.NewScanner(os.Stdin)
CLOSE:
	for sc.Scan() {
		args := strings.Fields(sc.Text())
		if len(args) == 0 {
			fmt.Fprintf(conn, "empty request\n")
			continue
		}
		cmd := args[0]
		switch cmd {
		case "close":
			fmt.Fprintln(conn, sc.Text())
			break CLOSE
		case "cd", "ls", "get":
			fmt.Fprintln(conn, sc.Text())
		default:
			fmt.Fprintf(conn, "%s is not supported\n", cmd)
			continue
		}
	}

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
