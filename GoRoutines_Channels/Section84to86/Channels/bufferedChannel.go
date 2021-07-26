package main

import (
	"fmt"
	"time"
)

func main() {
	res := mirroredQuery()
	fmt.Println(res)
}

func mirroredQuery() string {
	var responses = make(chan string)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	return <-responses // return the quickest response
}
func request(hostname string) (response string) {
	if hostname == "asia.gopl.io" {
		time.Sleep(1 * time.Second)
		return hostname
	} else if hostname == "europe.gopl.io" {
		time.Sleep(1 * time.Second)
		return hostname
	} else {
		time.Sleep(1 * time.Second)
		return hostname
	}
}
