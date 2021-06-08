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

func BitPopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func RightShiftPopCount(x uint64) int {
	var ret int
	for i := 0; i < 64; i++ {
		if x&(1<<i) != 0 {
			ret += 1
		}
	}
	return ret
}

func PopCountClear1(x uint64) int {
	retval := 0
	for x > 0 {
		x = x & (x - 1)
		retval++
	}
	return retval
}

func main() {
	fmt.Println(PopCount(63))
}

func PopCount(x uint64) int {
	var ret byte
	for i := 0; i < 8; i++ {
		ret += pc[byte(x>>(i*8))]
	}
	return int(ret)
}
