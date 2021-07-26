package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	//create a listener first
	list, l_err := net.Listen("tcp", "localhost:8000")
	if l_err != nil {
		log.Fatal(l_err)
	}

	//Now, accept the connections
	for {
		conn, c_err := list.Accept()
		if c_err != nil {
			log.Println(c_err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	sc := bufio.NewScanner(c)
	for sc.Scan() {
		text := sc.Text()
		go echo(c, text)
	}
	c.Close()
}

func echo(c net.Conn, text string) {
	fmt.Println("\t", strings.ToUpper(text))
	time.Sleep(1 * time.Second)
	fmt.Println("\t", text)
	time.Sleep(1 * time.Second)
	fmt.Println("\t", strings.ToLower(text))
	time.Sleep(1 * time.Second)
}
