package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "All the $glitters are not $gold"
	fmt.Println(s)
	expand(s, foo)
}

func expand(s string, f func(string) string) {
	words := strings.Split(s, " ")
	for i, v := range words {
		if strings.HasPrefix(v, "$") {
			words[i] = foo(v[1:])
		}
	}

	fmt.Println(strings.Join(words, " "))
}

func foo(s string) string {
	return s
}
