// package main

// import (
// 	"io"
// 	"log"
// 	"net"
// 	"os"
// )

// func main() {
// 	conn, err := net.Dial("tcp", "localhost:8000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	done := make(chan struct{})
// 	c, ok := conn.(*net.TCPConn)
// 	if !ok {
// 		log.Fatal(err)
// 	}

// 	go func() {

// 		io.Copy(os.Stdout, c)
// 		log.Println("done")
// 		done <- struct{}{} // メインゴルーチンへ通知
// 	}()
// 	mustCopy(c, os.Stdin)
// 	c.CloseWrite()
// 	// c.CloseRead()
// 	<-done
// }

// func mustCopy(dst io.Writer, src io.Reader) {
// 	_, err := io.Copy(dst, src)
// 	if err == io.EOF {
// 		return
// 	}
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.CloseWrite()
	<-done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		if err == io.EOF {
			return
		}
		log.Fatal(err)
	}
}
