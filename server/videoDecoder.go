package main

import (
	"log"

	"github.com/giorgisio/goav/avcodec"
)

func createVideoDecoder() *avcodec.Context {
	codec, err := avcodec.FindDecoder(avcodec.CodecId(avcodec.AV_CODEC_ID_H264))
	if err != nil {
		log.Fatalf("Failed to find decoder: %v", err)
	}

	decoderCtx := avcodec.AvcodecAllocContext3(codec)
	if decoderCtx.AvcodecOpen2(codec, nil) < 0 {
		log.Fatalf("Failed to open decoder")
	}

	return decoderCtx
}
