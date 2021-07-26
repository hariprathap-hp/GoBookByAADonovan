package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	w.Write([]byte("Hariprathap\n"))
	fmt.Println("hari")
}
