package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // Например, обрыв соединения
			continue
		}
		go handleConn(conn)
		log.Printf("client connected to FTP on: %s", conn.LocalAddr())
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	sc := bufio.NewScanner(conn)
	cwd := "."
CLOSE:
	for sc.Scan() {
		args := strings.Fields(sc.Text())
		cmd := args[0]
		switch cmd {
		case "close":
			fmt.Fprintln(conn, "close command:")
			break CLOSE
		case "ls":
			if len(args) < 2 {
				fmt.Fprintln(conn, "ls command:")
				err := ls(conn, cwd)
				if err != nil {
					fmt.Fprint(conn, err)
				}
			} else {
				path := args[1]
				if err := ls(conn, path); err != nil {
					fmt.Fprint(conn, err)
				}
			}
		case "cd":
			if len(args) < 2 {
				fmt.Fprintf(conn, "not enough arguments\n")
			} else {
				if _, err := os.ReadDir(cwd + "/" + args[1]); err != nil {
					fmt.Fprintln(conn, err)
					continue
				}
				cwd += "/" + args[1]
			}
		case "get":
			if len(args) < 2 {
				fmt.Fprintf(conn, "not enough arguments\n")
			} else {
				file, err := os.ReadFile(args[1])
				if err != nil {
					fmt.Fprintf(conn, "read file err: %v", err)
					log.Printf("read file err: %v", err)
				}
				fmt.Fprintf(conn, "%s\n", file)
			}
		case "send":
			if err := writeFile(conn, args, sc); err != nil {
				fmt.Fprintf(conn, "send file error: %v", err)
			}
		}
	}
	if sc.Err() != nil {
		fmt.Fprintln(conn, sc.Err().Error())
	}
}

func writeFile(w io.Writer, args []string, sc *bufio.Scanner) error {
	filename := args[1]
	file, err := os.Create("__" + filename)
	if err != nil {
		fmt.Fprint(w, err)
		return err
	}
	defer func(f *os.File) {
		f.Close()
		fmt.Fprintf(w, "%s sent. filename=%q\n", filename, f.Name())
	}(file)

	c, err := strconv.Atoi(args[2])
	if err != nil {
		log.Println(err)
	}
	var data string
	for i := 0; i < c && sc.Scan(); i++ {
		data += sc.Text() + "\n"
	}

	fmt.Fprintln(file, strings.TrimSuffix(data, "\n\n"))
	log.Printf("get file: %q", file.Name())
	return sc.Err()
}

func ls(w io.Writer, dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	printDir(files, w)
	return nil
}

func printDir(files []os.DirEntry, writer io.Writer) {
	const format = "%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(writer, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Name", "Mode", "Length")
	fmt.Fprintf(tw, format, "----", "----", "------")
	dir := "dir"
	for _, file := range files {
		if !file.IsDir() {
			dir = "file"
		}
		fmt.Fprintf(tw, format, file.Name(), file.Type(), dir)
	}
	tw.Flush()
}
