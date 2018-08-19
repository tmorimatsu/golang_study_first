package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	var data []string
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex03: %v\n", err)
		os.Exit(1)
	}
	for _, d := range getTextNodes(data, doc) {
		if d != "" {
			fmt.Println(d)
		}
	}
}

func getTextNodes(data []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		data = append(data, n.Data)
	}

	c := n.FirstChild
	if c != nil && n.Data != "script" && n.Data != "style" {
		data = getTextNodes(data, c)
	}
	c = n.NextSibling
	if c != nil {
		data = getTextNodes(data, c)
	}

	return data
}
