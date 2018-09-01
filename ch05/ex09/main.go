package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "I $dont have any money I am laterally broke"
	f := func(s string) string { return "do" }
	fmt.Println(expand(s, f))
}

func expand(s string, f func(string) string) string {
	if !strings.Contains(s, "$") {
		return s
	}

	words := strings.Split(s, " ")
	for i, word := range words {
		if !strings.HasPrefix(word, "$") {
			continue
		}

		words[i] = f(strings.TrimLeft(word, "$"))
	}
	return strings.Join(words, " ")
}
