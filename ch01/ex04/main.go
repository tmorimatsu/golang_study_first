package main

import (
	"bufio"
	"fmt"
	"os"
)

// 重複した行のそれぞれが含まれていた全てのファイルの名前を出力する用に修正しなさい

func main() {
	fileCounts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, fileCounts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, fileCounts)
			f.Close()
		}
	}
	for line, n := range fileCounts {
		fmt.Println(line)
		for filename, i := range n {
			fmt.Printf("%d\t%s\n", i, filename)
		}
		fmt.Println("")
	}
}

func countLines(f *os.File, fileCounts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if fileCounts[input.Text()] == nil {
			fileCounts[input.Text()] = make(map[string]int)
		}
		fileCounts[input.Text()][f.Name()]++
	}
}
