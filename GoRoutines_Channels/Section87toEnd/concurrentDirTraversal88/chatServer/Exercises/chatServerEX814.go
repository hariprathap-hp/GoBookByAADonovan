package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	go broadcaster()
	http.HandleFunc("/", chatHandler)
	http.ListenAndServe("localhost:8000", nil)
}

//create a type for a string channel
type client chan<- string

var (
	messages = make(chan string)
	entering = make(chan client)
	leaving  = make(chan client)
)

func broadcaster() {
	//create a slice of channels
	var clients = make(map[client]bool)

	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}

}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside chatHandler")
	if err := r.ParseForm(); err != nil {
		log.Print("Parse Error", err)
	}
	var who string

	for k, v := range r.Form {
		if k == "name" {
			who = v[0]
			break
		}
	}

	if who == "" {
		who = r.RemoteAddr
	}
	hi, ok := w.(http.Hijacker)
	if !ok {
		log.Fatalln("Can't Hijack.")
	}
	conn, _, err := hi.Hijack()
	if err != nil {
		log.Fatalln("Hijack error")
	}

	//create a channel which will be updated whenever a new client comes
	var ch = make(chan string)
	//whenever a new client enters, update the channel for the same so it will be broadcasted
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
