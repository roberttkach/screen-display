package main

import (
	"github.com/giorgisio/goav/avcodec"
)

func createVideoDecoder() (*avcodec.Context, error) {
	codec, err := avcodec.FindDecoder(avcodec.CodecId(avcodec.AV_CODEC_ID_H264))
	if err != nil {
		return nil, err
	}

	decoderCtx := avcodec.AvcodecAllocContext3(codec)
	if decoderCtx.AvcodecOpen2(codec, nil) < 0 {
		return nil, err
	}

	return decoderCtx, nil
}
