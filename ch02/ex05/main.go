package main

import (
	"fmt"
	"os"
	"strconv"

	"./popcount"
)

/*
パフォーマンスを比較する
*/

func main() {
	num, _ := strconv.Atoi(os.Args[1])
	fmt.Println(popcount.PopCount(uint64(num)))
	fmt.Println(popcount.PopCountClearLowestBit(uint64(num)))
}
