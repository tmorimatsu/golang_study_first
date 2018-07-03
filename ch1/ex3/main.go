package main

import (
	"fmt"
	"os"
	"time"

	"./joinargs"
)

// リファクタとefficient.go, inefficient.goの削除

func main() {
	fmt.Println("Start effcient method")
	start := time.Now()
	joinargs.Efficient(os.Args[1:])
	nanoSecs := time.Since(start).Nanoseconds()
	fmt.Println(nanoSecs)

	fmt.Println("Start ineffcient method")
	start = time.Now()
	joinargs.Inefficient(os.Args[1:])
	nanoSecs = time.Since(start).Nanoseconds()
	fmt.Println(nanoSecs)
}
