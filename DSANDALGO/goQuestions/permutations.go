package main

import (
	"fmt"
	"time"
)

var res = []string{}

func main() {
	str := "cabin"
	start := time.Now()
	permutation(str, 0)
	fmt.Println(time.Since(start).Seconds())
	/*i := 1
	for _, v := range res {
		if len(v) == len(str) {
			fmt.Println(i, v)
			i += 1
		}
	}*/
}

func permutation(s string, pos int) {
	if len(s) == pos {
		return
	}
	permutation(s, pos+1)
	perm(pos, s)
}

func perm(pos int, s string) {
	if len(res) == 0 {
		res = append(res, s[pos:])
		return
	}

	for _, v := range res {
		for a := 0; a <= len(v); a += 1 {
			if a == 0 {
				r := s[pos:pos+1] + v[a:]
				res = append(res, r)
			} else {
				r := v[:a] + s[pos:pos+1] + v[a:]
				res = append(res, r)
			}

		}
	}

}
