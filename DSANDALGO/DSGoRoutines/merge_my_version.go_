package main

import "fmt"

func main() {
	//declare your slice
	var arr = []int{26, 5, 2021, 29, 7, 31, 59, 111, 92}
	fmt.Println("Array before sorting is ")
	print(arr)
	l := len(arr)
	mergeSort(arr, 0, l-1)
	fmt.Println("Array after sorting is ")
	print(arr)
}

func print(arr []int) {
	for i, v := range arr {
		fmt.Printf("%d - %d\n", i, v)
	}
}

func mergeSort(arr []int, s, e int) {
	//return if the length of the arr is 1 or less than 1
	if s >= e {
		return
	}
	m := (s + e) / 2
	mergeSort(arr, s, m)
	mergeSort(arr, m+1, e)
	merge(arr, s, m, e)

}

func merge(a []int, s, m, e int) {
	final := []int{}
	//create two arrays that contain the copy of the original array elements
	a1 := a[s:m]
	a2 := a[m:e]
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)
	i := 0
	j := 0

	for i < len(a1) && j < len(a2) {
		if a1[i] <= a2[j] {
			final = append(final, a1[i])
			i += 1
		} else {
			final = append(final, a2[j])
			j += 1
		}
	}
	for i < len(a1) {
		final = append(final, a1[i])
		i += 1
	}
	for j < len(a2) {
		final = append(final, a2[j])
		j += 1
	}

	fmt.Println("Final is -- ", final)
	for s < e {
		a[s] = final[s]
		s += 1
	}
	fmt.Println("a is -- ", a)
}
