// tmpconvパッケージは摂氏(Celsius) と華氏(Fahrenheit) の温度変換を行います
package main

import (
	"fmt"

	"./tempconv"
)

// TODO: パッケージのテストを書く

func main() {
	fmt.Println(tempconv.CToF(100))
	fmt.Println(tempconv.FToC(100))
	fmt.Println(tempconv.CToK(100))
}
