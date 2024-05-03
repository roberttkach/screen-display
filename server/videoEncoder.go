package main

import (
	"image"
	"log"

	"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avutil"
)

func encodeVideo(img image.Image, width int, height int, fps int) (*avutil.Frame, error) {
	codec, err := avcodec.FindEncoder(avcodec.CodecId(avcodec.AV_CODEC_ID_H264))
	if err != nil {
		log.Fatalf("Failed to find encoder: %v", err)
	}

	encoderCtx := avcodec.AvcodecAllocContext3(codec)
	if encoderCtx.AvcodecOpen2(codec, nil) < 0 {
		log.Fatalf("Failed to open encoder")
	}

	frame := avutil.AvFrameAlloc()
	frame.SetWidth(width)
	frame.SetHeight(height)
	frame.SetFormat(int(avutil.AV_PIX_FMT_YUV420P))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			y := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			u := -0.14713*float64(r) - 0.28886*float64(g) + 0.436*float64(b)
			v := 0.615*float64(r) - 0.51498*float64(g) - 0.10001*float64(b)
			frame.Data()[y*frame.Linesize()[0]+x] = uint8(y)
			frame.Data()[y*frame.Linesize()[1]+x/2] = uint8(u)
			frame.Data()[y*frame.Linesize()[2]+x/2] = uint8(v)
		}
	}

	if encoderCtx.AvcodecEncodeVideo2(frame, nil) < 0 {
		log.Fatalf("Failed to encode frame")
	}

	return frame, nil
}
