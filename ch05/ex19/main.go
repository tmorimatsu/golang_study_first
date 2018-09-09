package main

import "fmt"

func main() {
	defer func() {
		p := recover()
		fmt.Println(p)
	}()
	nonReturn()
}

func nonReturn() {
	panic("test panic")
}
