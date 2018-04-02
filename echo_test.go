package main

import (
	"bytes"
	"net"
	"testing"
)

func assertEqual(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Fatalf("%v != %v", actual, expected)
	}
}

func TestMain(t *testing.T) {
	go startTCPServer()
	go startUDPServer()
}

func TestTCPEcho(t *testing.T) {
	message := "hello, 世界\n"

	conn, err := net.Dial("tcp", ":8080")
	assertEqual(t, err, nil)

	defer conn.Close()

	n, err := conn.Write([]byte(message))
	assertEqual(t, err, nil)

	assertEqual(t, n, len(message))

	b := make([]byte, len(message)*2)
	n, err = conn.Read(b)
	assertEqual(t, err, nil)
	assertEqual(t, n, len(message))

	b = b[:n]
	assertEqual(t, bytes.Equal(b, []byte(message)), true)
}

func TestUDPEcho(t *testing.T) {
	message := "hello, 世界\n"

	conn, err := net.Dial("udp", ":8080")
	assertEqual(t, err, nil)

	defer conn.Close()

	n, err := conn.Write([]byte(message))
	assertEqual(t, err, nil)

	assertEqual(t, n, len(message))

	b := make([]byte, len(message)*2)
	n, err = conn.Read(b)
	assertEqual(t, err, nil)
	assertEqual(t, n, len(message))

	b = b[:n]
	assertEqual(t, bytes.Equal(b, []byte(message)), true)
}
