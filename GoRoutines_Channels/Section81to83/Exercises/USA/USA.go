package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	listener, c_err := net.Listen("tcp", "localhost:8010")
	if c_err != nil {
		fmt.Println(c_err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	loc, _ := time.LoadLocation("US/Eastern")
	io.WriteString(c, time.Now().In(loc).Format("15:04:05\n"))
}
