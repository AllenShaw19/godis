package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net"
)

func ListenAndServe(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("listen err: %v", err)
	}
	defer listener.Close()

	log.Printf("bind: %s, start listening...", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("accept err: %v", err)
		}
		go Handle(conn)
	}
}

func Handle(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Printf("conn close")
			} else {
				log.Printf("read from conn err: %v", err)
			}
			return
		}
		b := []byte(msg)
		conn.Write(b)
	}
}

func main() {
	ListenAndServe(":8088")
}
