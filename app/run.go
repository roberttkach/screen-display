package main

import (
	"log"
	"net"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run(conn net.Conn) {
	cfg := pixelgl.WindowConfig{
		Title:  "Screen Share",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	decoder, err := createVideoDecoder()
	if err != nil {
		log.Fatalf("Failed to create decoder: %v", err)
		return
	}

	for !win.Closed() {
		win.Clear(colornames.Black)

		frame, err := captureScreen()
		if err != nil {
			log.Fatalf("Failed to capture screen: %v", err)
			return
		}

		pkt, err := encodeVideo(frame, 1280, 720, 30)
		if err != nil {
			log.Fatalf("Failed to encode video: %v", err)
			return
		}

		if err := sendVideo(pkt, conn); err != nil {
			log.Fatalf("Failed to send video: %v", err)
			return
		}

		videoData := <-dataChan

		frame, err = decoder.Decode(videoData)
		if err != nil {
			log.Fatalf("Failed to decode video frame: %v", err)
			return
		}

		img := frameToImage(frame)

		picData := pixel.PictureDataFromImage(img)
		pic := pixel.NewSprite(picData, picData.Bounds())

		pic.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

		win.Update()
	}
}
