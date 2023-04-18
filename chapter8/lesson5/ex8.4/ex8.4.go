/*
Exercise 8.4
Модифицируйте сервер reverb2 так, чтобы он использовал по одному
объекту sync.WaitGroup для каждого соединения для подсчета количества активных
горутин echo. Когда он обнуляется, закрывайте пишущую половину
TCP-соединения, как описано в упражнении 8.3. Убедитесь, что вы изменили
клиентскую программу netcat3 из этого упражнения так, чтобы она ожидала
последние ответы от параллельных горутин сервера даже после закрытия
стандартного ввода.
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
	var wg sync.WaitGroup

	input := bufio.NewScanner(conn)
	for input.Scan() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			echo(conn, input.Text(), 1*time.Second)
		}()

	}
	wg.Wait()
	conn.Close()
}

func echo(c io.Writer, text string, duration time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(text))
	time.Sleep(duration)
	fmt.Fprintln(c, "\t", text)
	time.Sleep(duration)
	fmt.Fprintln(c, "\t", strings.ToLower(text))
}
