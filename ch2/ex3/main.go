package main

import (
	"fmt"
	"os"
	"strconv"

	"./popcount"
)

/*
単一の式を使う代わりにループを使う
パフォーマンスを比較する
*/

func main() {
	num, _ := strconv.Atoi(os.Args[1])
	fmt.Println(popcount.PopCount(uint64(num)))
	fmt.Println(popcount.PopCountLoop(uint64(num)))
}
