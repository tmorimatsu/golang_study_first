package popcount

import "fmt"

var pc [256]byte

func init() {
	// 8bitの取りうる全ての値のpopulation countを準備している
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	fmt.Println(byte(256 >> (1 * 8)))
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	// 練習問題2.3
	var num int
	for i := 0; i < 8; i++ {
		num += int(pc[byte(x>>(uint64(i)*8))])
	}
	return num
}

func PopCountBitShift64(x uint64) int {
	// 練習問題2.4
	var num int
	for i := 0; i < 8; i++ {
		num += int(pc[byte(x>>(uint64(i)*8))])
	}
	return num
}

func PopCountClearLowestBit(x uint64) int {
	// 練習問題2.5
	var num int
	for i := 0; i < 8; i++ {
		num += int(pc[byte(x>>(uint64(i)*8))])
	}
	return num
}
