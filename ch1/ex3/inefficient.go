package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var s string
	for _, arg := range os.Args[1:] {
		s +=  arg + " "
	}
	nanoSecs := time.Since(start).Nanoseconds()
	fmt.Println(nanoSecs)
}