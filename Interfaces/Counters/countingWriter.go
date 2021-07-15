package main

import (
	"fmt"
	"io"
	"os"
)

type Wraps struct {
	bc io.Writer
	nb int64
}

func (wrp *Wraps) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	n, _ := wrp.bc.Write(p)
	wrp.nb += int64(n)
	return n, nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := Wraps{
		bc: w,
		nb: 0,
	}
	return &cw, &cw.nb
}

func main() {
	fd, _ := os.Create("/home/hari//go/src/TheGoProgLangBook/Interfaces/h_file")
	w, n := CountingWriter(fd)
	w.Write([]byte("Hariprathap"))
	fmt.Println(w, *n)
}
