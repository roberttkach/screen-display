package main

import (
	"net"
	"testing"

	"github.com/3d0c/gmf"
)

func TestCaptureScreen(t *testing.T) {
	// Capture Screen
	frame, err := captureScreen()
	if err != nil {
		t.Errorf("Failed to capture screen: %v", err)
	}
	if frame == nil {
		t.Errorf("Failed to capture screen: frame is nil")
	}
}

func TestCreateVideoDecoder(t *testing.T) {
	// Creating a decoder
	decoder, err := createVideoDecoder()
	if err != nil {
		t.Errorf("Failed to create video decoder: %v", err)
	}
	if decoder == nil {
		t.Errorf("Failed to create video decoder: decoder is nil")
	}
}

func TestEncodeVideo(t *testing.T) {
	// Encoding
	frame := gmf.NewFrame()
	pkt, err := encodeVideo(frame, 1280, 720, 30)
	if err != nil {
		t.Errorf("Failed to encode video: %v", err)
	}
	if pkt == nil {
		t.Errorf("Failed to encode video: packet is nil")
	}
}

func TestFrameToImage(t *testing.T) {
	// Convert frame to image
	frame := gmf.NewFrame()
	img := frameToImage(frame)
	if img == nil {
		t.Errorf("Failed to convert frame to image")
	}
}

func TestListenConnection(t *testing.T) {
	// Listen connection
	conn, _ := net.Dial("tcp", "localhost:8000")
	go listenConnection(conn)
}

func TestMain(t *testing.T) {
	// Start the main function
	go main()
}

func TestRun(t *testing.T) {
	// Run run
	conn, _ := net.Dial("tcp", "localhost:8000")
	go run(conn)
}

func TestSendVideo(t *testing.T) {
	// Send video
	conn, _ := net.Dial("tcp", "localhost:8000")
	pkt := gmf.NewPacket()
	err := sendVideo(pkt, conn)
	if err != nil {
		t.Errorf("Failed to send video: %v", err)
	}
}
