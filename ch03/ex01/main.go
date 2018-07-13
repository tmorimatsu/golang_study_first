package main

import (
	"fmt"
)

/*
関数fが有限ではないfloat64値を返すならば、SVGファイルは不正な<Polygon>要素を含むことになります。
(もっとも、多くのSVGレンダラはそれをうまく処理しますが)
不正なポリゴンをスキップするようにプログラムを修正しなさい
*/

func main() {
	fmt.Println("test")
}

// 不正なポリゴンかどうかを判定するメソッド
func ValidateInfinityFloat() bool {
	return false
}
