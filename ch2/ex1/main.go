// tmpconvパッケージは摂氏(Celsius) と華氏(Fahrenheit) の温度変換を行います
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"./tempconv"
)

/*
 TODO:
 パッケージのテストを書く
*/

func main() {

	tmp, err := strconv.Atoi(os.Args[1])
	unit := ""
	if len(os.Args) > 2 {
		unit = os.Args[2]
		// Celsiusとかに変換してString()してしまえば自動で単位つくのでは？
		fmt.Printf("%d °%s: \n", tmp, strings.ToUpper(unit))

	}
	if err != nil {
		fmt.Println(err)
	}
	switch unit {
	case "c":
		fmt.Println(" " + tempconv.CToF(tempconv.Celsius(tmp)).String())
		fmt.Println(" " + tempconv.CToK(tempconv.Celsius(tmp)).String())
	case "f":
		fmt.Println(" " + tempconv.FToC(tempconv.Fahrenheit(tmp)).String())
		fmt.Println(" " + tempconv.FToK(tempconv.Fahrenheit(tmp)).String())
	case "k":
		fmt.Println(" " + tempconv.KToC(tempconv.Kelvin(tmp)).String())
		fmt.Println(" " + tempconv.KToF(tempconv.Kelvin(tmp)).String())
	default:
		fmt.Println("第二引数に単位(c, f, k のいずれか)を指定してください")
	}
}
