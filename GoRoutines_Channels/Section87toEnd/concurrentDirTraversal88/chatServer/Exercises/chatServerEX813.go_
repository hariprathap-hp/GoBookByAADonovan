package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
	}

	go broadCaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConn(conn)
	}
}

type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadCaster() {
	fmt.Printf("Type of entering channel is -- %T\n", entering)
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			fmt.Println("messages")
			for cli := range clients {
				fmt.Printf("%s\n", msg)
				cli <- msg + " is here dude"
			}
		case cli := <-entering:
			fmt.Printf("%T\n", cli)
			fmt.Println("Entering")
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}

}

func handleConn(conn net.Conn) {
	//create a channel which will be updated whenever a new client comes
	var ch = make(chan string)
	//whenever a new client enters, update the channel for the same so it will be broadcasted
	who := conn.RemoteAddr().String()
	go clientWriter(conn, ch, who)
	ch <- "you are " + who
	messages <- who + " is here"
	entering <- ch

	timeout := time.NewTimer(5 * time.Second)
	go func() {
		<-timeout.C
		conn.Close()
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		timeout.Reset(5 * time.Second)
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}
func clientWriter(c net.Conn, ch <-chan string, who string) {
	fmt.Println(who, "Client Writer")
	for msg := range ch {
		fmt.Fprintln(c, msg)
	}
}
