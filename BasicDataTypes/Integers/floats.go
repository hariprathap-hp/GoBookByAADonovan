package main

import "fmt"

func main() {
	var x = 1.23456
	fmt.Printf("%f\n%[1]g\n%[1]e\n", x)

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)
}
