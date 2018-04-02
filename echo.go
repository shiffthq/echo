package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"unicode/utf8"
)

const defaultListenPort = 8080

func startTCPServer() {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(defaultListenPort))
	if err != nil {
		log.Fatalf("listen fail, err=%s", err)
	}

	defer ln.Close()

	log.Printf("listen tcp on %s", ln.Addr())

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("accept fail, err=%s", err)
		}

		log.Printf("accept connection from: %v", conn.RemoteAddr())
		go handleTCPConnection(conn)
	}
}

func handleTCPConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr()
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		b, err := reader.ReadBytes('\n')
		if err != nil {
			switch err {
			case io.EOF:
				log.Printf("(%s) remote closed", remoteAddr)
			}
			if err != io.EOF {
				log.Printf("(%s) read fail, err=%s", remoteAddr, err)
			}
			conn.Close()
			break
		}
		log.Printf("(%s): %d/%s", remoteAddr, utf8.RuneCountInString(string(b)), string(b))
		conn.Write(b)
	}
}

func startUDPServer() {
	addr, err := net.ResolveUDPAddr("udp", ":"+strconv.Itoa(defaultListenPort))
	if err != nil {
		log.Fatalf("ResolveUDPAddr fail, err=%s", err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("listen fail, err=%s", err)
	}

	log.Printf("listen udp on %s", conn.LocalAddr())

	for {
		handleUDPConnection(conn)
	}
}

func handleUDPConnection(conn *net.UDPConn) {
	b := make([]byte, 1500)

	n, remoteAddr, err := conn.ReadFrom(b)
	if err != nil {
		log.Fatalf("ReadFrom fail, err=%v", err)
		return
	}

	b = b[:n]
	log.Printf("(%v): %d/%s", remoteAddr, utf8.RuneCountInString(string(b)), string(b))
	_, err = conn.WriteTo(b, remoteAddr)
	if err != nil {
		log.Fatalf("WriteTo fail, err=%v", err)
	}
}

func main() {
	log.Printf("Start with pid: %d", os.Getpid())

	go startTCPServer()
	go startUDPServer()

	ch := make(chan error)
	<-ch
}
