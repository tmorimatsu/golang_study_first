package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	counter := make(map[string]int)
	counter["a"] = 0
	counter["p"] = 0
	counter["div"] = 0
	counter["span"] = 0
	counter["other"] = 0

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex02: %v\n", err)
		os.Exit(1)
	}

	for key := range countNodesByElement(counter, doc) {
		fmt.Printf("key:%s , value:%d\n", key, counter[key])
	}
}

func countNodesByElement(counter map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			counter["a"]++
		case "p":
			counter["p"]++
		case "div":
			counter["div"]++
		case "span":
			counter["span"]++
		default:
			counter["other"]++
		}
	}

	c := n.FirstChild
	if c != nil {
		counter = countNodesByElement(counter, c)
	}
	c = n.NextSibling
	if c != nil {
		counter = countNodesByElement(counter, c)
	}

	return counter
}
