package main

import (
	"fmt"
	"unicode"
)

/**
* UTF-8エンコードされた[]byteスライス内で隣接している
* Unicodeスペース(unicode.IsSpaceを参照)を、
* もとのスライス内で一つのASCIIスペースへ圧縮する関数を書きなさい。
 */

func main() {
	b := []byte("t    etet   e")
	fmt.Println(compressSpaces(b))
}

func compressSpaces(b []byte) string {
	for i := 0; i < len(b); i++ {
		if unicode.IsSpace(rune(b[i])) && b[i] == b[i+1] {
			b = append(b[:i], b[i+1:]...)
			i--
		}
	}
	return string(b)
}
