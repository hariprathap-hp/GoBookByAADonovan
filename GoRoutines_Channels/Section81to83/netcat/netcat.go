package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, c_err := net.Dial("tcp", "localhost:8000")
	if c_err != nil {
		log.Println(c_err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src net.Conn) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
