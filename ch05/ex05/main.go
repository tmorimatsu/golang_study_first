package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("args must be a url")
		os.Exit(1)
	}
	words, images, err := CountWordsAndImage(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("words: %d, images: %d", words, images)
}

func CountWordsAndImage(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
	}
	var wordsCount, imgsCount int
	words, images = countWordsAndImages(wordsCount, imgsCount, doc)
	return words, images, nil
}

// HTMLドキュメント内に含まれる単語と画像の数を返します。
func countWordsAndImages(wordsCount, imgsCount int, n *html.Node) (words, images int) {

	// 単語数をカウント
	if n.Type == html.TextNode {
		wordsCount += wordCount(n.Data)
	}

	// 画像の数をカウント
	if n.Type == html.ElementNode && n.Data == "img" {
		imgsCount++
	}

	c := n.FirstChild
	if c != nil {
		wordsCount, imgsCount = countWordsAndImages(wordsCount, imgsCount, c)
	}
	c = n.NextSibling
	if c != nil {
		wordsCount, imgsCount = countWordsAndImages(wordsCount, imgsCount, c)
	}

	return wordsCount, imgsCount
}

// 与えられた文字列の単語数を返却します
func wordCount(s string) int {
	return len(strings.Fields(s))
}
