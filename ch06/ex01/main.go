package main

import (
	"bytes"
	"fmt"
)

// TODO: テスト

func main() {
	var test IntSet
	test.words = []uint64{1, 2, 3, 10, 14}
	test.Len()
	fmt.Printf("test len : %v\n", test)
	test.Remove(10)
	fmt.Printf("test remove : %v\n", test)
	a := test.Copy()
	fmt.Printf("a : %v\n", *a)
	fmt.Printf("test copy : %v\n", test)
	test.Clear()
	fmt.Printf("test clear : %v\n", test)

	test.AddAll(6, 5, 4, 3)
	fmt.Printf("test addall : %v\n", test)

	var t IntSet
	t.words = []uint64{11,3,3,4,5,6,7,8}
	test.IntersectWith(&t)
	fmt.Printf("test intersectwith : %v\n", test)

	t.words = []uint64{6,5,0,3}
	test.DifferenceWith(&t)
	fmt.Printf("test differencewith : %v\n", test)

	t.words = []uint64{6,5,11,3}
	test.SymmetricDifference(&t)
	fmt.Printf("test symmetricdifference : %v\n", test)

}

type IntSet struct {
	words []uint64
}

// Hasは負ではない値xをセットに含んでいるか否かを報告します。
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Addはセットに負ではない値xを追加します。
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWithは、sとtの和集合をsに設定します。
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String は"{1 2 3}"の形式の文字列としてセットを返します。
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('{')
	return buf.String()
}


// ex01
func (s *IntSet) Len() int {
	return len(s.words)
}

func (s *IntSet) Remove(x int) {
	for i, val := range s.words {
		if val == uint64(x) {
			s.words = append(s.words[:i], s.words[i+1:]...)
			return
		}
	}
}

func (s *IntSet) Clear() {
	s.words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
	var i IntSet
	copyWords := make([]uint64, len(s.words))
	copy(copyWords, s.words)
	i.words = copyWords
	return &i
}


// ex02
func (s *IntSet) AddAll(a ...int) {
	au := make([]uint64, len(s.words))
	for _, val := range a {
		au = append(au, uint64(val))
	}
	s.words = append(s.words, au...)
}


// ex03
// IntersectWithは、sとtの共通部分をsに設定します。
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

// DifferenceWithは、sに対するtの差集合(s-t)をsに設定します。
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

// SymmetricDifferenceは、sとtの対象集合をsに設定します。
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
			s.words[i] &^= (s.words[i] & tword)
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// ex04
// TODO: 問題文の理解
func (s *IntSet) Elems(t *IntSet) {

}