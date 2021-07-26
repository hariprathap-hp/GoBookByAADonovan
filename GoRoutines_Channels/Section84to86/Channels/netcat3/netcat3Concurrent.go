package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", "localhost:8000")
	c, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(err)
	}
	done := make(chan struct{})

	defer func() {
		c.CloseWrite()
		<-done
	}()
	go func() {
		io.Copy(os.Stdout, c)
		fmt.Println("done")
		c.CloseRead()
		done <- struct{}{}
	}()

	mustCopy(c, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	fmt.Println("Inside MustCopy")
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Outside MustCopy")
}
