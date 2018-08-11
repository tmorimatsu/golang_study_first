package main

import (
	"strings"
)

// TODO: マルチバイト対応

func main() {
	println(isAnagram("guamunara", "anaguramu"))
	println(isAnagram("アナグラム", "グアムナラ"))
	println(isAnagram("testing", "ingtset"))
	println(isAnagram("testing", "iangtset"))
	println(isAnagram("testing", "inttgst"))
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		idx := strings.Index(string(s2), string(s1[i]))
		if idx < 0 {
			return false
		}
		s2 = s2[0:idx] + s2[idx+1:]
	}
	return true
}
