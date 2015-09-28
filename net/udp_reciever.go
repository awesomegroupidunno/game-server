package net

import (
	"github.com/awesomegroupidunno/game-server/encoder"
	"log"
	"net"
)

const default_max_packet = 8192
const default_port = ":10001"

type UdpReceiver struct {
	PortNumber     string
	MaxPacket      int
	EncoderDecoder encoder.EncoderDecoder
}

func (u *UdpReceiver) Run() {

	if u.MaxPacket == 0 {
		u.MaxPacket = default_max_packet
	}

	if u.PortNumber == "" {
		u.PortNumber = default_port
	}

	if u.EncoderDecoder == nil {
		u.EncoderDecoder = &encoder.JsonEncoderDecoder{Tag: "Default"}
	}

	go u.start()
}

func (u *UdpReceiver) start() {
	server_address, err_add := net.ResolveUDPAddr("udp", u.PortNumber)

	if err_add != nil {
		log.Fatal(err_add)
	}

	connection, err_con := net.ListenUDP("udp", server_address)

	if err_con != nil {
		log.Fatal(err_con)
	}

	for {
		u.handleUdp(connection)
	}

}

func (u *UdpReceiver) handleUdp(conn *net.UDPConn) {
	buffer := make([]byte, u.MaxPacket)

	n, address, readError := conn.ReadFromUDP(buffer)
	a, encode_error := u.EncoderDecoder.Decode(buffer[:n])

	if readError == nil && encode_error == nil {
		log.Println(a)
		conn.WriteToUDP([]byte("Ack"), address)
	} else {
		log.Println(readError)
		log.Println(encode_error)
		log.Println(address)
	}

}
