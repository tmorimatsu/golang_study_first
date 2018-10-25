package main

import (
	"sync"
)

var initLoader sync.Once
var pc [256]byte

func main() {
	PopCount(100)
}

// 関数名がinitであると呼び出しも参照もできず、プログラムの開始時点で1回だけ呼ばれる
func initLoad() {
	// 8bitの取りうる全ての値のpopulation countを準備している
	for i := range pc {
		// 右シフトして最小の桁を確認
		pc[i] = pc[i/2] + byte(i&1)
	}
}

var t int

func PopCount(x uint64) int {
	initLoader.Do(initLoad)
	num := int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
	t += num
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
