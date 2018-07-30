package main

import (
	"fmt"
)

func main() {
	arr := []string{"test", "tete1", "tete2", "tete2", "tete2", "tete2", "tete3", "test", "test"}
	t := nonContiguousValues(arr)
	fmt.Println(t)
}

func nonContiguousValues(arr []string) []string {
	num := len(arr)
	a := 1
	for i := 0; i < num-a; i++ {
		if arr[i] == arr[i+1] {
			arr = append(arr[:i], arr[i+1:]...)
			a++
			i--
		}
	}
	return arr
}
