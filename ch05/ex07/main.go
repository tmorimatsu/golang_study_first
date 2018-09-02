package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex07: %v\n", err)
		os.Exit(1)
	}

	forEachNode(doc, startElement, endElement)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {

	// 子要素がない場合
	if n.FirstChild == nil && n.Type == html.ElementNode {
		fmt.Printf("%*s<%s/>\n", depth*2, "", prettyElement(n))
		return
	}

	if n.Type == html.ElementNode {
		element := fmt.Sprintf("%*s<%s>\n", depth*2, "", prettyElement(n))
		fmt.Print(element)
		depth++
		return
	}

	if n.Type == html.TextNode {
		text := n.Data

		// 改行コードを削除します
		text = strings.TrimSpace(text)
		for strings.HasPrefix(text, "\n") || strings.HasSuffix(text, "\n") {
			text = strings.TrimLeft(text, "\n")
			text = strings.TrimRight(text, "\n")
		}

		text = prettyText(text, fmt.Sprintf("%*s", depth*2, ""))

		if strings.TrimSpace(text) == "" {
			return
		}

		if !strings.HasSuffix(text, "\n") {
			text = text + "\n"
		}
		fmt.Printf("%s", text)
		return
	}

	if n.Type == html.CommentNode {
		comment := fmt.Sprintf("%*s<!-- %s -->\n", depth*2, "", n.Data)
		fmt.Print(comment)
		return
	}
}

func endElement(n *html.Node) {

	// 子要素がない場合
	if n.FirstChild == nil {
		return
	}

	if n.Type == html.ElementNode {
		depth--
		element := fmt.Sprintf("%*s</%s>\n", depth*2, "", n.Data)
		fmt.Print(element)
		return
	}

	if n.Type == html.TextNode || n.Type == html.CommentNode {
		return
	}
}

func prettyElement(n *html.Node) string {
	attr := " "
	for _, a := range n.Attr {
		attr = attr + a.Key + "=\"" + a.Val + "\""
	}
	return strings.TrimRight(fmt.Sprintf("%s%s", n.Data, attr), " ")
}

// 改行コード後にprefixをつけます
func prettyText(s string, prefix string) string {
	texts := strings.Split(s, "\n")
	for i, text := range texts {
		texts[i] = prefix + text
	}
	s = strings.Join(texts, "\n")
	return s
}
