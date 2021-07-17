package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://espncricinfo.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "http Get failed with error %v\n", err)
		os.Exit(1)
	}

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Reading from response body failed with error %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	fmt.Printf("%s", byt)
}
