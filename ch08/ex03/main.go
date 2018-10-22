package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	c, ok := conn.(*net.TCPConn)
	if !ok {
		log.Fatal(err)
	}

	go func() {

		io.Copy(os.Stdout, c)
		log.Println("done")
		done <- struct{}{} // メインゴルーチンへ通知
	}()
	mustCopy(c, os.Stdin)
	c.CloseWrite()
	// c.CloseRead()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err == io.EOF {
		return
	}
	if err != nil {
		log.Fatal(err)
	}
}
