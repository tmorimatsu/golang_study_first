package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(rotate(arr[:], 2))
}

func rotate(arr []int, i int) []int {
	tmp := arr[:i]
	arr = arr[i:]
	return append(arr, tmp...)
}
