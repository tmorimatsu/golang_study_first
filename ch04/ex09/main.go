package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	wordfreq("./sample.txt")
}

func wordfreq(s string) {
	frequency := make(map[string]int)
	file, err := os.Open(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fialed to open file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReaderSize(file, 4096)

	for {
		r, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("ファイル読み込みエラー: %v", err)
			os.Exit(1)
		}
		words := strings.Fields(string(r))
		for _, w := range words {
			word := w
			frequency[word]++
		}
	}

	for word, n := range frequency {
		fmt.Printf("%-30s %d\n", word, n)
	}
}
