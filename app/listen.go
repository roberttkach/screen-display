package main

import (
	"log"
	"net"
)

func listenConnection(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatalf("Failed to read: %v", err)
			return
		}
		dataChan <- buf[:n]
	}
}
