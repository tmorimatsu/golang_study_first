package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	filename, n, err := fetch(os.Args[1])
	fmt.Printf("%s | %d | %v", filename, n, err)
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	n, err = io.Copy(f, resp.Body)
	//// 根本原因の可能性が高いio.Copyのエラーを優先して返却する
	//if closeErr := f.Close(); err == nil {
	//	err = closeErr
	//}
	return local, n, err
}
