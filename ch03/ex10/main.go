package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Println(comma(os.Args[1]))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(s[i : i+1])
		if i%3 == n%3-1 && i != n-1 {
			buf.WriteString(",")
		}
	}
	return buf.String()
}
