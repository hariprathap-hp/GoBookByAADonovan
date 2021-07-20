package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	Words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < 64 && s.Words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.Words) {
		s.Words = append(s.Words, 0)
	}
	s.Words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.Words {
		if i < len(s.Words) {
			s.Words[i] |= tword
		} else {
			s.Words = append(s.Words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.Words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte('-')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var i IntSet
	i.Add(5)
	i.Add(6)
	i.Add(69)
	i.Add(70)
	for i, v := range i.Words {
		fmt.Println(i, v)
	}
	fmt.Println(i.Has(42))
	fmt.Println(i.String())
}
