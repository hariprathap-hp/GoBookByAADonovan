package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//create a reader to reader input from io console
	reader := bufio.NewReader(os.Stdin)

	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(char)

	switch char {
	case 'A':
		fmt.Println("A is pressed")
		//fallthrough
	case 'B':
		fmt.Println("B is pressed")
		//fallthrough
	default:
		fmt.Println("Another key is pressed")
	}
}
