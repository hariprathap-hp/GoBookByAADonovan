package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
	}

	for {
		conn, c_err := listener.Accept()
		if c_err != nil {
			fmt.Println(c_err)
		}

		go handleFunc(conn)
	}
}

func handleFunc(c net.Conn) {
	fmt.Println("Inside Handle Function")
	defer c.Close()
	//first read info from the connection using scanner
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		cmds := strings.Split(scanner.Text(), " ")
		fmt.Println(cmds)
		switch cmds[0] {
		case "ls":
			exeCommand(c, cmds[0], cmds[1:]...)
		case "cd":
			fmt.Println(len(cmds))
			if len(cmds) > 1 {
				fmt.Println("first argument", cmds[1])
				if err := os.Chdir(cmds[1]); err != nil {
					log.Print(err)
				}
			} else {
				fmt.Println("Please provide an argument to cd, now moving to root directory")
				if err := os.Chdir("/"); err != nil {
					log.Print(err)
				}
			}
		case "get":
			file, err := os.Open(cmds[1])
			if err != nil {
				log.Printf("file %s: %v", cmds[1], err)
				continue
			}
			mustCopy(c, file)
		case "close":
			return
		default:
			help := "ls: list content\ncd: change directory\nget: get file content\n" +
				"close: close connection\n"
			mustCopy(c, strings.NewReader(help))
		}
	}
}
func mustCopy(w io.Writer, r io.Reader) {
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
}

func exeCommand(w io.Writer, e string, args ...string) {
	cmd := exec.Command(e, args...)
	cmd.Stdout = w
	if err := cmd.Run(); err != nil {
		log.Print(err)
	}
}
