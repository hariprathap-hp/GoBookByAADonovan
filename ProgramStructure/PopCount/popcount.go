package main

import "fmt"

//declare a byte array which stores 256 byte values
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		//fmt.Printf("i is -- %v, pc[i/2] -- %v , byte(i&1) --  %v , result -- %v\n", i, pc[i/2], byte(i&1), pc[i])
	}
}

func main() {
	fmt.Println(PopCount(63))
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
