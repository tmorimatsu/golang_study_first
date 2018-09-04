package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// (WIP)

func main() {
	breadthFirst(crawl, os.Args[1:])
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	i := 0
	for len(worklist) > 0 || i > 100 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
				i++
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, path, content, err := extract(url)
	if err != nil {
		log.Fatal(err)
	}
	err = cloneHtml(path, content)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

func cloneHtml(path, content string) error {

	os.MkdirAll(path, 0777)
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	f, err := os.OpenFile(path+"index.html", os.O_WRONLY|os.O_CREATE, 0777)
	defer f.Close()
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)
	fmt.Fprint(w, content)
	return nil
}

func extract(url string) (linkList []string, path string, content string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, "", "", err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, "", "", fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	// ioutil.ReadAll(resp.Body) をすると後のParseに影響が出る
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	content = string(byteArray)
	path = "./" + resp.Request.URL.Host + resp.Request.URL.Path

	doc, err := html.Parse(strings.NewReader(content))
	resp.Body.Close()
	if err != nil {
		return nil, "", "", fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		fmt.Println(n.Data)
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	fmt.Println(links)
	return links, path, content, nil
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
