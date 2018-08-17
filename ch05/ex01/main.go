package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	for _, link := range body(os.Stdin) {
		fmt.Println(link)
	}
}

func body(f *os.File) []string {
	doc, err := html.Parse(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	return visit(nil, doc)
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	c := n.FirstChild
	if c != nil {
		links = visit(links, c)
	}
	c = n.NextSibling
	if c != nil {
		links = visit(links, c)
	}

	return links
}
