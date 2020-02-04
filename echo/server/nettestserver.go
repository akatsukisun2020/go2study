package main

import (
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Println(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go echoFunc(conn)
	}
}

func echoFunc(conn net.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}
		// conn.Write(buf[:n])
		log.Println(string(buf[:n]))
	}
}
