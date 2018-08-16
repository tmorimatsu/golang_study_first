package main

import "fmt"

func main() {
	n := []int{5,12,4,6,23,66,100,033}
	fmt.Println(n)
	n = Sort(n)
	fmt.Println(n)
}

type tree struct {
	value int
	left, right *tree
}

func Sort(values []int) []int {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	return values
}

// appendValues は t の要素をvalues の正しい順序に追加し、結果のスライスを返します。
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// return &tree{value: value} と同じ
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}