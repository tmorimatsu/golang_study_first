package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex04: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, n.Data+": "+a.Val)
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
