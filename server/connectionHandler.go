package main

import (
	"log"
	"sync"

	"github.com/hashicorp/yamux"
)

func handleConnection(server *yamux.Server, mu *sync.Mutex, activeConnections *int) {
	defer func() {
		mu.Lock()
		(*activeConnections)--
		mu.Unlock()
	}()

	for {
		stream, err := server.AcceptStream()
		if err != nil {
			log.Printf("Failed to accept stream: %v", err)
			return
		}

		go func() {
			buf := make([]byte, 1024)
			for {
				n, err := stream.Read(buf)
				if err != nil {
					log.Printf("Failed to read from stream: %v", err)
					return
				}

				decoder := createVideoDecoder()
				frame, err := decoder.Decode(buf[:n])
				if err != nil {
					log.Printf("Failed to decode video frame: %v", err)
					return
				}

				img := frameToImage(frame)

				dataChan <- img

				imgToSend := <-dataChan

				frameToSend, err := encodeVideo(imgToSend, 1280, 720, 30)
				if err != nil {
					log.Printf("Failed to encode video: %v", err)
					return
				}

				if _, err := stream.Write(frameToSend.Data()); err != nil {
					log.Printf("Failed to write to stream: %v", err)
				}
			}
		}()
	}
}
