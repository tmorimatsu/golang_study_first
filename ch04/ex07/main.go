package main

import "fmt"

/**
* UTF-8でエンコードされた文字列を表す[]byteスライス文字を、
* そのスライス内で逆順にするようにreverseを修正しなさい。
* 新たなメモリを割り当てることなく行えるでしょうか？
*
 */

 WIP

func main() {
	b := []byte("abcdefgh")
	fmt.Println(string(reverse(b)))
}

func reverse(b []byte) []byte {

	if len(b) < 2 {
		return b
	}

	b = append(reverse(), reverse()...)

	return b
}
