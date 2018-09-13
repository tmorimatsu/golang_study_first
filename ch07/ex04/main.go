package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
)

func main() {
	var r NewReader
	a := r.SetString(`<html><a>tttete</a></html>`)
	doc, err := html.Parse(a)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		links = append(links, n.Data)

	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

type NewReader struct {
	src    string
	offset int
}

// TODO: fixme
func (r NewReader) Read(p []byte) (n int, err error) {
	n = 0
	for {
		if r.offset+n >= len(r.src) {
			r.offset += n
			return n, io.EOF
		}
		if n >= len(p) {
			r.offset += n
			return n, nil
		}
		p[n] = r.src[r.offset+n]
		n++
	}
}

func (r NewReader) SetString(s string) io.Reader {
	r.src = s
	return r
}
