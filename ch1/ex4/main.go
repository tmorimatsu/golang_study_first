package main

import (
	"bufio"
	"fmt"
	"os"
)

// 重複した行のそれぞれが含まれていた全てのファイルの名前を出力する用に修正しなさい

// FIXME: 動かない時がある

func main() {
	file_counts :=  make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, file_counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, file_counts)
			f.Close()
		}
	}
	for line, n :=  range file_counts {
		fmt.Println(line)
		for filename, i :=  range n {
				fmt.Printf("%d\t%s\n", i, filename)
		}
		fmt.Println("")
	}
}

func countLines(f *os.File, file_counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if file_counts[input.Text()] == nil {
			file_counts[input.Text()] = make(map[string]int)
		}
		file_counts[input.Text()][f.Name()]++
	}
}
