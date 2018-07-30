package main

func main() {

}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// 拡大する余地のある。スライスを拡張する。
		z = x[:zlen]
	} else {
		// 十分な領域がない。新たな配列を割り当てる。
		// 計算量を線形に均すために倍に拡張する。
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
