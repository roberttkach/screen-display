package main

import (
	"github.com/3d0c/gmf"
)

func encodeVideo(frame *gmf.Frame, width int, height int, fps int) (*gmf.Packet, error) {
	codec, err := gmf.FindEncoder("libx264")
	if err != nil {
		return nil, err
	}

	cc := gmf.NewCodecCtx(codec)
	defer gmf.Release(cc)

	cc.SetWidth(width).SetHeight(height)
	cc.SetTimeBase(gmf.AVR{Num: 1, Den: fps})
	cc.SetPixFmt(gmf.AV_PIX_FMT_YUV420P).SetProfile(gmf.FF_PROFILE_H264_HIGH)

	if err := cc.Open(nil); err != nil {
		return nil, err
	}

	pkt := gmf.NewPacket()
	if pkt, err = cc.Encode(frame, pkt); err != nil {
		return nil, err
	}

	return pkt, nil
}
