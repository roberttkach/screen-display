package main

import (
	"bytes"
	"image/png"

	"github.com/3d0c/gmf"
	"github.com/kbinani/screenshot"
)

func captureScreen() (*gmf.Frame, error) {
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil, err
	}

	frame, err := gmf.NewFrame().SetBuffer(buf.Bytes())
	if err != nil {
		return nil, err
	}

	return frame, nil
}
