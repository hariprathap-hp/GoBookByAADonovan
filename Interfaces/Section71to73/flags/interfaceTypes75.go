package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	var s io.Writer

	//In the below statement, Stdout is the concrete type and s in the interface type
	s = os.Stdout
	fmt.Println(s.Write([]byte("Hari")))

	var x interface{} = time.Now()
	fmt.Println(x)
}
