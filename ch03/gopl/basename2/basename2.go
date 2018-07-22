package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := os.Args[1]
	fmt.Println(basename(s))
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // "/"が見つからなければ-1
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
