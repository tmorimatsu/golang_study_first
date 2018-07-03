package main

import (
	"fmt"
	"os"
	"strconv"

	"./popcount"
)

/*
TODO:
パフォーマンスを比較する
*/

func main() {
	num, _ := strconv.Atoi(os.Args[1])
	fmt.Println(popcount.PopCount(uint64(num)))
	fmt.Println(popcount.PopCountLoop(uint64(num)))
}
