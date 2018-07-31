package main

import (
	"fmt"
)

/**
* UTF-8エンコードされた[]byteスライス内で隣接している
* Unicodeスペース(unicode.IsSpaceを参照)を、
* もとのスライス内で一つのASCIIスペースへ圧縮する関数を書きなさい。
 */

func main() {
	b := []byte("tetet e    ")
	compressSpaces(b)
}

func compressSpaces(b []byte) {
	for i := 0; i < len(b)-1; i++ {
		fmt.Println(b[i])
		if b[i] == []byte(" ")[0] && b[i] == b[i+1] {
			// c := unicode.IsSpace(rune(b[i]))
			// fmt.Println(c)
		}
	}
}
