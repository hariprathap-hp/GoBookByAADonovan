package main

import (
	"fmt"
	"sort"
)

func main() {
	myList := []int{3, 8, 1, 9, 12, 15, 21, 99, 43, 54}
	sort.Ints(myList)
	binarySearch(43, 0, 9, myList)
}

func binarySearch(k int, start, end int, list []int) {
	mid := (end + start) / 2

	if mid <= end && mid >= start {
		if list[mid] == k {
			fmt.Printf("Element %d found at %dth position\n", k, mid)
			return
		} else if k < list[mid] {
			binarySearch(k, start, mid-1, list)
		} else {
			binarySearch(k, mid+1, end, list)
		}
	}
}
