package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	counter := make(map[string]int)

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex02: %v\n", err)
		os.Exit(1)
	}

	for key := range countNodesByElement(counter, doc) {
		fmt.Printf("%s: %d\n", key, counter[key])
	}
}

func countNodesByElement(counter map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		counter[n.Data]++
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
