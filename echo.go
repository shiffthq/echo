package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

const defaultListenPort = 8080

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()
	defer log.Printf("Connection closed: %s", conn.RemoteAddr())

	log.Println("Accept connection from:", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	for {
		b, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				log.Println("Read fail:", err.Error())
			}
			break
		}
		log.Printf("Received (%s): %s", conn.RemoteAddr(), b)
		conn.Write(b)
	}
}

func main() {
	log.Printf("Start with pid: %d", os.Getpid())

	ln, err := net.Listen("tcp", ":"+strconv.Itoa(defaultListenPort))
	if err != nil {
		panic("Listen fail: " + err.Error())
	}
	log.Printf("Listen %s", ln.Addr())

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf(err.Error())
			panic("Listen fail:" + err.Error())
		}
		go handleTCPConnection(conn)
	}
}
