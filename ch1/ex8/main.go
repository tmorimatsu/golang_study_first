package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// http://がなければ追加するように修正。 strings.HasPrefixを使う

func main() {
	for _, url := range os.Args[1:] {
		// todo:別関数で定義してtestを作る
		originalUrl := url
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		resp, err :=  http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		resp.Body.Close()
		if originalUrl != url {
			fmt.Println("original url: " + originalUrl)
			fmt.Println("fixed url:    " + url)
		} else {
			fmt.Println("url: " + url)
		}
	}
}