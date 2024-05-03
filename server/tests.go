package main

import (
	"crypto/tls"
	"image"
	"image/color"
	"net"
	"sync"
	"testing"

	"github.com/giorgisio/goav/avutil"
	"github.com/hashicorp/yamux"
)

func TestMain(t *testing.T) {
	// Load certificates
	_, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		t.Errorf("Failed to load keys: %v", err)
	}
}

func TestFrameToImage(t *testing.T) {
	// Convert frame to image
	frame := avutil.Frame{}
	img := frameToImage(&frame)
	if img == nil {
		t.Errorf("Failed to convert frame to image")
	}
}

func TestHandleConnection(t *testing.T) {
	// Processing the connection
	conn, _ := net.Dial("tcp", "localhost:8000")
	server, _ := yamux.Server(conn, nil)
	var mu sync.Mutex
	activeConnections := 0
	go handleConnection(server, &mu, &activeConnections)
}

func TestCreateVideoDecoder(t *testing.T) {
	// Decoder for video
	decoder := createVideoDecoder()
	if decoder == nil {
		t.Errorf("Failed to create video decoder")
	}
}

func TestEncodeVideo(t *testing.T) {
	// Video encoding
	img := image.NewRGBA(image.Rect(0, 0, 1280, 720))
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			img.Set(x, y, color.RGBA{uint8(x % 256), uint8(y % 256), 0, 255})
		}
	}
	frame, err := encodeVideo(img, 1280, 720, 30)
	if err != nil {
		t.Errorf("Failed to encode video: %v", err)
	}
	if frame == nil {
		t.Errorf("Failed to encode video: frame is nil")
	}
}
