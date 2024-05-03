package main

import (
	"net"

	"github.com/3d0c/gmf"
)

func sendVideo(pkt *gmf.Packet, conn net.Conn) error {
	if _, err := conn.Write(pkt.Data()); err != nil {
		return err
	}
	return nil
}
