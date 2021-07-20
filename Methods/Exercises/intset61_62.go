package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (i *IntSet) Has(u uint64) bool {
	word, byte := u/64, u%64
	isPresent := false
	if word < 64 {
		if 1<<byte&i.words[word] != 0 {
			isPresent = true
		}
	}
	return isPresent
}

func (i *IntSet) Add(u int) {
	//first the word number in the slice of words in which this value has to be inserted
	word, byte := u/64, uint(u%64)
	for word >= len(i.words) {
		i.words = append(i.words, 0)
	}
	i.words[word] |= 1 << byte
}

func (i *IntSet) String() string {
	var res bytes.Buffer

	//check bit by bit of each word and print it in string format
	//first find the length of the words
	if len(i.words) == 0 {
		return res.String()
	}
	res.WriteByte('{')
	for n, word := range i.words {
		if word == 0 {
			continue
		}
		//for each word, iterate 64 bytes and if any byte is set and add that number to th string
		for i := 0; i < 64; i++ {
			if (1<<uint(i))&word != 0 {
				if res.Len() > len("{") {
					res.WriteByte(' ')
				}
				fmt.Fprintf(&res, "%d", 64*n+i)
			}
		}
	}
	res.WriteByte('}')
	return res.String()
}

func (i *IntSet) Len() int {
	if len(i.words) == 0 {
		return 0
	}

	var res int
	for _, word := range i.words {
		if word == 0 {
			continue
		}
		//for each word, iterate 64 bytes and if any byte is set and add that number to th string
		for i := 0; i < 64; i++ {
			if (1<<uint(i))&word != 0 {
				res++
			}
		}
	}
	return res
}

func (i *IntSet) Remove(x int) {
	if len(i.words) == 0 {
		return
	}
	//find the word number and the bit
	word, bit := x/64, x%64
	i.words[word] ^= 1 << bit
}

func (i *IntSet) Clear() {
	for n, word := range i.words {
		if word == 0 {
			continue
		}
		i.words[n] = 0
	}
}

func (i *IntSet) Copy() *IntSet {
	//var newSet *IntSet
	newSet := new(IntSet)
	newSet.words = make([]uint64, len(i.words))
	copy(newSet.words, i.words)
	return newSet

}

func (i *IntSet) AddAll(n ...int) {
	for _, v := range n {
		byte_num, bit := v/64, uint(v%64)
		for byte_num >= len(i.words) {
			i.words = append(i.words, 0)
		}
		i.words[byte_num] |= 1 << bit
	}
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) InterSectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words[i] &= 0
		}
	}
}
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			break
		}
	}
}

func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			break
		}
	}
}

func main() {
	var x IntSet
	x.Add(34)
	x.Add(35)
	x.Add(78)
	x.Add(114)
	x.Add(33333)
	x.Add(223)
	fmt.Println(x.String())
	fmt.Println("Length is ", x.Len())
	x.Remove(223)
	fmt.Println(x.String())
	//x.Clear()
	newSet := x.Copy()
	fmt.Println("NewSet", newSet.String())
	x.AddAll(1, 2, 3, 4, 5, 6)
	fmt.Println(x.String())

	var r1, r2 IntSet
	r1.AddAll(1, 2, 3, 4, 5, 6, 7, 76)
	r2.AddAll(3, 4, 5, 76, 84, 85)
	fmt.Println("r1", r1.String())
	fmt.Println("r2", r2.String())
	//r1.InterSectWith(&r2)
	//fmt.Println("r1 Intersection", r1.String())
	r1.DifferenceWith(&r2)
	fmt.Println("r1 Difference with r2", r1.String())
	//r1.SymmetricDifferenceWith(&r2)
	//fmt.Println("r1 Symmteric Difference with r2", r1.String())
}
