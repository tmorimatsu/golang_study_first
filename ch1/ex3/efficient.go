package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	//fmt.Println(strings.Join(os.Args[1:], " "))
	strings.Join(os.Args[1:], " ")
	nanoSecs := time.Since(start).Nanoseconds()
	fmt.Println(nanoSecs)
}