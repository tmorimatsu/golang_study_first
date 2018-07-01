package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"./conv"
)

/*
TODO:
(必須)
- []メソッドのテストを書く
(余裕があれば)
- []変換された値を丸める
- []全角数字の対応
*/

func main() {
	if len(os.Args) > 2 {
		fmt.Println("引数は一つだけにしてください")
	} else if len(os.Args) == 2 {
		num, err := strconv.ParseFloat(os.Args[1], 64)
		if err != nil {
			fmt.Println("数値を入力してください")
			os.Exit(1)
		}
		printAllConvertedUnit(num)
	} else {
		fmt.Println("変換する数値を入力してください")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("数値を入力してください")
			os.Exit(1)
		}
		printAllConvertedUnit(num)
	}

}

func printAllConvertedUnit(num float64) {
	fmt.Println(strconv.FormatFloat(num, 'g', 6, 64) + "M -> " + conv.MeterToFeet(conv.Meter(num)).String())
	fmt.Println(strconv.FormatFloat(num, 'g', 6, 64) + "F -> " + conv.FeetToMeter(conv.Feet(num)).String())
	fmt.Println(strconv.FormatFloat(num, 'g', 6, 64) + "kg -> " + conv.KilogramToPond(conv.Kilogram(num)).String())
	fmt.Println(strconv.FormatFloat(num, 'g', 6, 64) + "lb -> " + conv.PondToKilogram(conv.Pond(num)).String())
	fmt.Println(strconv.FormatFloat(num, 'g', 6, 64) + "°C -> " + conv.CelsiusToFahrenheit(conv.Celsius(num)).String())
	fmt.Println(strconv.FormatFloat(num, 'g', 6, 64) + "°F -> " + conv.FahrenheitToCelsius(conv.Fahrenheit(num)).String())
}
