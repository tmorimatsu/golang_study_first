package main

import (
	"fmt"
	"reflect"
)

func main() {
	int := "test"
	fmt.Println(int)
	fmt.Println(reflect.TypeOf(int))

	num := 3
	const num2 int = 3
	fmt.Println(num)
	fmt.Println(reflect.TypeOf(num))
}
