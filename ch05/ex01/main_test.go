package main

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"testing"
)

// もうちょっと綺麗にかきたい

func TestSample(t *testing.T) {
	t.Log("元の関数の出力結果と同じこと")

	sampleF, err := os.Open("./sample.html")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer sampleF.Close()
	links := body(sampleF)

	ansF, err := os.Open("./ans.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer ansF.Close()

	var ans []string
	scanner := bufio.NewScanner(ansF)
	for scanner.Scan() {
		line := scanner.Text()
		ans = append(ans, line)
	}

	expected := links
	actual := ans
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
