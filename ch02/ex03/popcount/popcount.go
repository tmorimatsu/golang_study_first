package popcount

var pc [256]byte

func init() {
	// 8bitの取りうる全ての値のpopulation countを準備している
	for i := range pc {
		// 右シフトして最小の桁を確認
		pc[i] = pc[i/2] + byte(i&1)
	}
}

var t int

func PopCount(x uint64) int {
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

func PopCountLoop(x uint64) int {
	// 練習問題2.3
	var num int
	for i := 0; i < 8; i++ {
		num += int(pc[byte(x>>(uint64(i)*8))])
	}
	t += num
	return num
}

func PopCountBitShift64(x uint64) int {
	// 練習問題2.4
	num := 0
	current := x
	for i := 0; i < 64; i++ {
		if byte(current&1) == 1 {
			num++
		}
		current = current >> 1
	}
	return num
}

func PopCountClearLowestBit(x uint64) int {
	// 練習問題2.5
	num := 0
	current := x
	for true {
		if current == 0 {
			break
		}
		current = current & (current - 1)
		num++
	}
	return num
}
