package main

import (
	"image"
	"image/color"

	"github.com/giorgisio/goav/avutil"
)

func frameToImage(frame *avutil.Frame) image.Image {
	width := frame.Width()
	height := frame.Height()
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := frame.Data()[y*frame.Linesize()[0]+x]
			img.SetRGBA(x, y, color.RGBA{pixel, pixel, pixel, 255})
		}
	}

	return img
}
