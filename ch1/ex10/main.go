package main

import (
	"fmt"
	"io"
	//"io/ioutil"
	"net/http"
	"os"
	"time"
)

// 大量のデータを生成するサイトを見つけ、fetchallを2回実行し、キャッシュされているか？・同じ内容を取得できているか？を確認、出力内容をファイルに保存

// todo: 毎回同じ内容を取得できているかどうかの確認

func main() {
	start := time.Now()
	filename := "out.txt"
	resp_filename := os.Args[1]
	ch := make(chan string)
	for _, url := range os.Args[2:] {
		go fetch(url, ch, resp_filename) // ゴルーチンを開始
	}
	for range os.Args[2:] {
		res := <-ch
		fmt.Println(res) // chチャネルから受信
		f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Fialed to open the file")
			os.Exit(1)
		}
		f.Write([]byte(res+"\n"))
		f.Close()
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, resp_filename string) {
	start := time.Now()
	resp_f, err := os.OpenFile(resp_filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer resp_f.Close()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // chチャネルへ送信
		return
	}

	nbytes, err := io.Copy(resp_f, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s %s", secs, nbytes, url, resp.Header["X-Cache"])
}
