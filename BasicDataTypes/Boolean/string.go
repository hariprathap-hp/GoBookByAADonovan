package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "H\va\vr\vi\vP\vr\va\vt\vh\va\vp\""
	fmt.Println(s)
	h := "\\xhh"
	fmt.Println(h)

	p := "harryHariPrathap.go"
	fmt.Println(strings.HasPrefix(p, "harry"))
	fmt.Println(strings.Contains(p, "Hari"))

	i := strings.LastIndex(p, ".")
	si := p[:i]
	fmt.Println(si)

	fmt.Println(strings.Count(s, "\v"))

	fmt.Println(strings.Fields(s))

}
