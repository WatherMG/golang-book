/*
Exercise 8.1.1 - clockwall

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
	"text/tabwriter"
	"time"
)

type server struct {
	name    string
	host    string
	message string
}

func (s *server) String() string {
	return fmt.Sprintf("%s:%s", s.name, s.host)
}

func parseServers(args []string) ([]*server, error) {
	var servers = make([]*server, 0, len(args))
	for _, arg := range args {
		n := strings.SplitN(arg, "=", 2)
		if len(n) != 2 {
			return nil, fmt.Errorf("invalid argument: %s", arg)
		}
		servers = append(servers, &server{n[0], n[1], ""})
	}
	return servers, nil
}

func (s *server) getMessages(conn io.ReadCloser) error {
	defer conn.Close()
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		s.message = sc.Text()
	}
	return sc.Err()
}

func main() {
	t := 1 * time.Second
	servers, err := parseServers(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range servers {
		c, err := s.connectToSever()
		if err != nil {
			log.Fatal(err)
		}
		go func(c io.ReadCloser, s *server) {
			if err := s.getMessages(c); err != nil {
				log.Fatal(err)
			}
		}(c, s)
	}
	printMessages(t, servers)
}

func (s *server) connectToSever() (io.ReadCloser, error) {
	conn, err := net.Dial("tcp", s.host)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Connected to %q\n", s.host)
	return conn, nil
}

func printMessages(delay time.Duration, connections []*server) {
	const format = "%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Name", "Server", "Message")
	fmt.Fprintf(tw, format, "----", "------", "-------")
	for {
		for _, s := range connections {
			fmt.Fprintf(tw, format, s.name, s.host, s.message)
		}
		fmt.Fprintf(tw, format, "----", "------", "-------")
		tw.Flush()
		time.Sleep(delay)
	}
}
