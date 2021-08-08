package main

import "testing"

func TestSum(t *testing.T) {
	mysum := 29
	res_sum := sum([]int{19, 10})
	if mysum != res_sum {
		t.Fail()
	}
}
