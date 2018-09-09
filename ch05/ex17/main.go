package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

// TODO微妙

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex12: %v\n", err)
		os.Exit(1)
	}

	//images := ElementsByTagName(doc, "img")
	a := ElementsByTagName(doc, "a")

	//fmt.Println(images)
	for _, node := range a {
		fmt.Printf("<%s>\n", node.Data)
	}
}

func ElementsByTagName(n *html.Node, names ...string) []*html.Node {

	var nodes []*html.Node
	if n.Type == html.ElementNode {
		for _, name := range names {
			if n.Data == name {
				nodes = append(nodes, n)
				break
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, ElementsByTagName(c, names...)...)
	}

	return nodes
}
