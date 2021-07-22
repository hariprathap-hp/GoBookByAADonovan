package main

import (
	"fmt"
	"io"
	"os"
)

type countWriter struct {
	Writer io.Writer
	count  int64
}

func (cw *countWriter) Write(b []byte) (int, error) {
	n, err := cw.Writer.Write(b)
	if err != nil {
		fmt.Println(err)
	}
	cw.count += int64(n)
	return n, err
}

//It has to return a new writer "N" that wraps the original writer "0" and a pointer to int64 that returns the number of bytes written to the new Writer
func CountingWriter(o io.Writer) (io.Writer, *int64) {
	bc := countWriter{
		Writer: o,
		count:  0,
	}

	return &bc, &bc.count
}

func main() {
	//CountingWriter returns address to a countWriter struct and the number of bytes in it.
	//So far, they are uninitialized

	//A file is created with file name h_file
	fd, err := os.Create("/home/hari/golib/src/deadpoet/The_Go_Prog_Language/h_file")
	if err != nil {
		fmt.Println(err)
	}
	//That file is passed to the CountingWriter function and the file descriptor will be assigned to the writer

	writer, count := CountingWriter(fd)
	fmt.Printf("%T\n", writer)
	fmt.Fprint(writer, "Hello World!!\n")
	fmt.Println(*count)
}
