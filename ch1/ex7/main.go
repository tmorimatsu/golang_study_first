package main

import (
	"fmt"
	//"io/ioutil"
	"io"
	"net/http"
	"os"
)

// io.Copyを利用する。 必ずエラー結果は検査するようにする

func main() {
	for _, url := range os.Args[1:] {
		resp, err :=  http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		resp.Body.Close()
	}
}
