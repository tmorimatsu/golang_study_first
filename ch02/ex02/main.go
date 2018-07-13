package main

import (
	"fmt"
	"os"
	"strconv"

	"./conv"
)

/*
TODO:
(必須)
- []変換された値を丸める
- []panicではなくerrorをreturnに含める
- []メソッドのテストを書く(WIP)
(余裕があれば)
- []全角数字の対応
- []引数がなかった時のscanner
*/

func main() {
	/*if len(os.Args) > 3 {
		fmt.Println("引数は 変換する数字 単位(温度:t, 長さ:l, 重さ:w) としてください")
	} else */if len(os.Args) == 3 {
		num, err := strconv.ParseFloat(os.Args[1], 64)
		if err != nil {
			fmt.Println("数値を入力してください")
			os.Exit(1)
		}
		unit := os.Args[2]
		switch unit {
		case "t":
			printConvertedTenparature(num)
		case "l":
			printConvertedLength(num)
		case "w":
			printConvertedWeigh(num)
		default:
			fmt.Println("引数は 変換する数字 単位(温度:t, 長さ:l, 重さ:w) としてください")
		}
	} else {
		/*fmt.Println("変換する数値を入力してください")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("数値を入力してください")
			os.Exit(1)
		}
		printAllConvertedUnit(num)*/
		fmt.Println("引数は 変換する数字 単位(温度:t, 長さ:l, 重さ:w) としてください")
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

func printConvertedTenparature(num float64) {
	fmt.Println(strconv.FormatFloat(num, 'g', 6, 64) + "°C -> " + conv.CelsiusToFahrenheit(conv.Celsius(num)).String())
	fmt.Println(strconv.FormatFloat(num, 'g', 6, 64) + "°F -> " + conv.FahrenheitToCelsius(conv.Fahrenheit(num)).String())
}

func printConvertedLength(num float64) {
	fmt.Println(strconv.FormatFloat(num, 'g', 6, 64) + "M -> " + conv.MeterToFeet(conv.Meter(num)).String())
	fmt.Println(strconv.FormatFloat(num, 'g', 6, 64) + "F -> " + conv.FeetToMeter(conv.Feet(num)).String())
}

func printConvertedWeigh(num float64) {
	fmt.Println(strconv.FormatFloat(num, 'g', 6, 64) + "kg -> " + conv.KilogramToPond(conv.Kilogram(num)).String())
	fmt.Println(strconv.FormatFloat(num, 'g', 6, 64) + "lb -> " + conv.PondToKilogram(conv.Pond(num)).String())
}
