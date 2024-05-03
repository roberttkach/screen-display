package main

import (
	"crypto/tls"
	"image"
	"log"
	"sync"

	"github.com/hashicorp/yamux"
)

var dataChan = make(chan image.Image)

func main() {
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: false}

	ln, err := tls.Listen("tcp", ":8000", &config)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer ln.Close()

	var mu sync.Mutex
	activeConnections := 0

	for {
		mu.Lock()
		if activeConnections >= 2 {
			mu.Unlock()
			continue
		}
		mu.Unlock()

		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		server, err := yamux.Server(conn, nil)
		if err != nil {
			log.Printf("Failed to start Yamux server: %v", err)
			continue
		}

		mu.Lock()
		activeConnections++
		mu.Unlock()

		go handleConnection(server, &mu, &activeConnections)
	}
}
