package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"golang.org/x/net/html"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("one argument is needed")
	}
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex04: %v\n", err)
		os.Exit(1)
	}

	id := os.Args[1]
	forEachNode(doc, ElementByID, nil, id)
	fmt.Println("done")
}

func forEachNode(n *html.Node, pre, post func(n *html.Node, id string) bool, id string) {
	if pre != nil {
		if !pre(n, id) {
			fmt.Printf("stopped because id %s is found\n", id)
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post, id)
	}

	if post != nil {
		if !post(n, id) {
			return
		}
	}
}

func ElementByID(doc *html.Node, id string) bool {
	for _, attr := range doc.Attr {
		if attr.Key == "id" && attr.Val == id {
			return false
		}
	}
	return true
}
