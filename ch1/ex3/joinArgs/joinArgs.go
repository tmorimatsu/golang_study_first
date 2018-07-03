package joinargs

import (
	"strings"
)

func Efficient(args []string) string {
	s := strings.Join(args, " ")
	return s
}

func Inefficient(args []string) string {
	var s string
	for _, arg := range args {
		s += arg + " "
	}
	return s
}
