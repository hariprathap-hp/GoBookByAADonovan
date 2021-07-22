package main

import (
	"fmt"
	"io"
	"strings"
)

type limitReader struct {
	r        io.Reader
	n, limit int
}

func (lr *limitReader) Read(p []byte) (int, error) {
	n, err := lr.r.Read(p[:lr.limit])
	lr.n += n
	if n > lr.limit {
		err = io.EOF
	}
	return n, err
}

func getReader(r io.Reader, n int) io.Reader {
	sm := limitReader{
		r:     r,
		limit: n,
	}
	return &sm
}

func main() {
	var b = make([]byte, 40)
	sReader := strings.NewReader("All the glitters are not gold. Be a Roman while in Rome")
	res := getReader(sReader, 40)
	fmt.Println(res.Read(b))
	fmt.Println(string(b))

}
