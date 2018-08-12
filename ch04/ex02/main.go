package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

// フラグ判定のメソッドを作成？

func main() {
	lenArgs := len(os.Args)
	if lenArgs < 2 || 3 < lenArgs {
		fmt.Println("文字列 変換用フラグ[optional] で指定してください")
		return
	} else {
		if lenArgs == 2 {
			converted := convertToSHA256(os.Args[1])
			fmt.Printf("%x", converted)
			return
		}
		if lenArgs == 3 {
			switch os.Args[2] {
			case "-384":
				// SHA384のハッシュを出力
				converted := convertToSHA384(os.Args[1])
				fmt.Printf("%x", converted)
			case "-512":
				// SHA512のハッシュを出力
				converted := convertToSHA512(os.Args[1])
				fmt.Printf("%x", converted)
			default:
				fmt.Println("フラグは '-384', '-512' の2種類です")
			}
		}
	}
}

func convertToSHA256(s string) [32]byte {
	return sha256.Sum256([]byte(s))
}

func convertToSHA384(s string) [48]byte {
	return sha512.Sum384([]byte(s))
}

func convertToSHA512(s string) [64]byte {
	return sha512.Sum512([]byte(s))
}
