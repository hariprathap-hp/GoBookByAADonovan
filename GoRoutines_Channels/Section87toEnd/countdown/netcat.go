package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, c_err := net.Dial("tcp", "localhost:8000")
	if c_err != nil {
		fmt.Println(c_err)
	}

	go func() {
		io.Copy(conn, os.Stdin)
	}()

	mustcopy(os.Stdout, conn)
}

func mustcopy(dst io.Writer, src io.Reader) {
	io.Copy(dst, src)
}
