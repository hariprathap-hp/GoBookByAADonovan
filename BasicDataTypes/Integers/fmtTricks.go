package main

import "fmt"

func main() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o \n", o)

	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)

	//How to print runes
	ascii := 'a'
	//unicode := ' D '
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii, ascii) // "97 a 'a'"
	//fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 D ' D '"
	fmt.Printf("%d %[1]q\n", newline) // "10 '\n'"
}
