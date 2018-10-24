package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
	members  = make(chan []string)
	member   []string
)

// 新たに到着したクライアントに対して現在のクライアントを知らせる
// [WIP]

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			m := <-members
			clients[cli] = true
			// TODO: ここで今いるメンバーを出力する
			for _, c := range m {
				// messages <- "current: " + c
				log.Println(c)
			}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who + "\n"
	messages <- who + " has arrived\n"
	entering <- ch
	member = append(member, who)
	// log.Println(member)
	members <- member

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text() + "\n"
	}

	leaving <- ch
	messages <- who + " has left\n"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintf(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}