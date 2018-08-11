package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(comma(os.Args[1]))
}

func comma(s string) string {
	var sign string
	var decimal string
	if strings.Contains(s, "+") || strings.Contains(s, "-") {
		sign = s[0:1]
		s = s[1:]
	}
	if strings.Contains(s, ".") {
		decimal = s[strings.Index(s, "."):]
		s = s[0:strings.Index(s, ".")]
	}

	n := len(s)

	if n <= 3 {
		return sign + s
	}

	tmp := comma(s[:n-3]) + "," + s[n-3:]
	if sign != "" {
		tmp = sign + tmp
	}
	if decimal != "" {
		tmp = tmp + decimal
	}

	return tmp
}
