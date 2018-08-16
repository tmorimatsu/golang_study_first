package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"os"
	"bufio"
	"unicode"
	"unicode/utf8"
	"strings"
)

type Comic struct {
	Num int
	Year, Month, Day string
	Title string
	Transcript string
	Alt string
	Img string
}

func getComic(n int) (Comic, error) {
	var comic Comic
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", n)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return comic, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return comic, fmt.Errorf("can't get comic %d: %s", n, resp.StatusCode)
	}
	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return comic, err
	}
	return comic, nil
}

type WordIndex map[string]map[int]bool
type NumIndex map[int]Comic

func indexComics(comics chan Comic) (WordIndex, NumIndex) {
	wordIndex := make(WordIndex)
	numIndex := make(NumIndex)
	for comic := range comics {
		numIndex[comic.Num] = comic
		scanner := bufio.NewScanner(strings.NewReader(comic.Transcript))
		scanner.Split(ScanWords)
		for scanner.Scan() {
			token := strings.ToLower(scanner.Text())
			if _, ok := wordIndex[token]; !ok {
				wordIndex[token][comic.Num] = true
			}
		}
	}
	return wordIndex, numIndex
}

func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	i := 0
	start := 0
	stop := 0
	for i < len(data) {
		r, size := utf8.DecodeRune(data[i:])
		i += size
		if unicode.IsLetter(r) {
			start = i - size
			break
		}
	}
	for i < len(data) {
		r, size := utf8.DecodeRune(data[i:])
		i += size
		if unicode.IsLetter(r) {
			stop = i - size
			break
		}
	}
	if stop > start {
		token = data[start:stop]
	}
	return i, token, nil
}

func search() {

}

func main() {
	comic, err := getComic(200)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	printComic(comic)
}


func printComic(comic Comic) {
	fmt.Println(comic.Num)
	fmt.Println()
	fmt.Println()
	fmt.Println(comic.Year)
	fmt.Println(comic.Month)
	fmt.Println(comic.Day)
	fmt.Println()
	fmt.Println()
	fmt.Println(comic.Title)
	fmt.Println()
	fmt.Println()
	fmt.Println(comic.Transcript)
	fmt.Println()
	fmt.Println()
	fmt.Println(comic.Alt)
	fmt.Println()
	fmt.Println()
	fmt.Println(comic.Img)
}