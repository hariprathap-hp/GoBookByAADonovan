package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	list, l_err := net.Listen("tcp", "localhost:8000")
	if l_err != nil {
		fmt.Println(l_err)
	}

	for {
		conn, c_err := list.Accept()
		if c_err != nil {
			fmt.Println(c_err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	//create a scanner to read from the connection
	scanner := bufio.NewScanner(c)
	text := make(chan string)
	var wg sync.WaitGroup

	go func() {
		for scanner.Scan() {
			text <- scanner.Text()
		}
		close(text)
	}()

	for {
		ticker := time.NewTicker(10 * time.Second)
		select {
		case <-ticker.C:
			ticker.Stop()
			c.Close()
			fmt.Println("disconnect silent client")
			return
		case t, ok := <-text:
			if ok {
				wg.Add(1)
				go func(text string) {
					defer wg.Done()
					fmt.Println("Echo echo")
					//io.Copy(c, os.Stdin)
					echo(c, text)
				}(t)
			} else {
				wg.Wait()
				c.Close()
				return
			}
		}
	}
}

func echo(c net.Conn, s string) {
	fmt.Fprintf(c, "\t%s\n", strings.ToUpper(s))
	time.Sleep(2 * time.Second)
	fmt.Fprintf(c, "\t%s\n", (s))
	time.Sleep(2 * time.Second)
	fmt.Fprintf(c, "\t%s\n", strings.ToLower(s))
	time.Sleep(2 * time.Second)
}
