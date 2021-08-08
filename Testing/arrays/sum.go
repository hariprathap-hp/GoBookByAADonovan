package main

import "fmt"

func main() {
	arr := []int{5, 7, 8, 9}
	sum := sum(arr)
	fmt.Println(sum)
}

func sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
