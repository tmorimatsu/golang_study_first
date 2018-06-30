// tmpcomvパッケージは摂氏(Celsius) と華氏(Fahrenheit) の温度変換を行います
package main

import (
	"fmt"

	"./tempcomv"
)

// TODO: パッケージのテストを書く

func main() {
	fmt.Println(tempcomv.CToF(100))
	fmt.Println(tempcomv.FToC(100))
	fmt.Println(tempcomv.CToK(100))
}
