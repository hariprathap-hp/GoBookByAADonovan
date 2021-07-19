package main

import (
	"fmt"
	"os"
)

func main() {
	var rmdirs []func()
	for _, d := range tempDir() {
		dir := d
		os.MkdirAll(dir, 0777)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}

	fmt.Printf("%T\n", rmdirs)
	for _, rmdir := range rmdirs {
		rmdir() // clean up
	}
}
func tempDir() []string {
	s := []string{"/home/hari/personal/dir1", "/home/hari/personal/dir2", "/home/hari/personal/dir3", "/home/hari/personal/dir1/dir11", "/home/hari/personal/dir1/dir22"}
	return s
}
