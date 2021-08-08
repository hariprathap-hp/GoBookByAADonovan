package main

import (
	"fmt"
	"testing"
)

var cc = make(map[rune]int)

func TestPassCharcount(t *testing.T) {
	r := []rune{1, 2, '$', 4, 5, '%', '*', 1, 1, 1, 1, 2, 3, 2, 2, 2}
	for _, v := range r {
		cc[v]++
	}

	res := charcount(r)
	exp := true

	if len(cc) != len(res) {
		t.Error()
	}
	for i, _ := range res {
		if cc[i] != res[i] {
			exp = false
		}
	}
	if exp != true {
		t.Fail()
	}
}

func TestFailCharcount(t *testing.T) {
	r := []rune{1, 2, '$', 4, 5, '%', '*', 1, 1, 1, 1, 2, 3, 2, 2, 2}
	for _, v := range r {
		cc[v]++
	}
	cc[8]++

	res := charcount(r)
	exp := true

	if len(cc) == len(res) {
		t.Error()
	}
	for i, _ := range res {
		if cc[i] != res[i] {
			fmt.Printf("val %v\n", i)
			exp = true
			break
		}
	}
	if exp != true {
		t.Fail()
	}
}
