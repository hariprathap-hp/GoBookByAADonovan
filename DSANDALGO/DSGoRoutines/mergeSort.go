package main

import (
	"fmt"
	"time"
)

//var res []int

func main() {
	//create a slice that needs to be sorted using merge sort
	var arr = []int{39, 33, 27, 10}
	fmt.Println("Before SOrting")
	print(arr)
	//res = make([]int, len(arr))
	start := time.Now()
	res := mergeSort(arr)
	fmt.Printf("%f\n", time.Since(start).Seconds())
	fmt.Println("After Sorting")
	print(res)
}

func print(a []int) {
	for i, v := range a {
		fmt.Println(i, v)
	}
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])

	return merge(first, second)
}
func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
