package main

import (
	"log"
	"net"

	"github.com/faiface/pixel/pixelgl"
	"github.com/hashicorp/yamux"
)

var dataChan = make(chan []byte)

func main() {
	conn, err := net.Dial("tcp", "133.133.133.133:8000")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
		return
	}
	defer conn.Close()

	client, err := yamux.Client(conn, nil)
	if err != nil {
		log.Fatalf("Failed to start Yamux client: %v", err)
		return
	}

	stream, err := client.OpenStream()
	if err != nil {
		log.Fatalf("Failed to open stream: %v", err)
		return
	}

	go listenConnection(stream)

	pixelgl.Run(func() {
		run(stream)
	})
}
