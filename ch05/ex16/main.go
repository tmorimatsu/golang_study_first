package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []string{"test1", "test2", "test3", "test4"}
	fmt.Println(strings.Join(a, ":"))

	fmt.Println(join(":", a...))
}

func join(sep string, a ...string) string {

	var b strings.Builder

	// TODO: はじめにbufferを全て確保する

	for _, s := range a {
		b.Grow(len(a))
		b.Write([]byte(s))
		if !(a[len(a)-1] == s) {
			b.Write([]byte(sep))
		}
	}

	return b.String()
}
