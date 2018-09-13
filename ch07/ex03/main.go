package main

import(
	"fmt"
)

func main() {

	var t *tree
	t = add(t, 4)
	t = add(t, 42)
	t = add(t, 24)
	t = add(t, 43)
	t = add(t, 11)

	t.String()

}


type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
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
		// Equivalent to return &tree{value: value}.
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


func (t *tree) String() {
	fmt.Println(print(t))
}

func print(t *tree) string {
	s := ""
	if t.left == nil && t.right == nil {
		s = fmt.Sprintf("%v", t.value)
	} else if t.left == nil {
		s = fmt.Sprintf("%v", t.value) + ", right {" + fmt.Sprintf("%v", print(t.right)) + "}"
	} else if t.right == nil {
		s = fmt.Sprintf("%v", t.value) + ", left {" + fmt.Sprintf("%v", print(t.left)) + "}"
	} else {
		s = fmt.Sprintf("%v", t.value) + ", left {" + fmt.Sprintf("%v", print(t.left)) + "}, right {" + fmt.Sprintf("%v", print(t.right)) + "}"
	}
	return s
}