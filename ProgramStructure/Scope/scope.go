package main

import (
	"fmt"
)

func main() {
	x := "Hari"

	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%s\n", string(x))
	}
}
