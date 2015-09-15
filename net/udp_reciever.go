package net

import (
	"fmt"
	"log"
	"net"
	"time"
)

const default_max_packet = 8192
const default_port = ":10001"

type UdpReceiver struct {
	PortNumber string
	MaxPacket  int
}

func (u *UdpReceiver) Run() {

	if u.MaxPacket == 0 {
		u.MaxPacket = default_max_packet
	}

	if u.PortNumber == "" {
		u.PortNumber = default_port
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
		buffer := make([]byte, u.MaxPacket)
		_, remoteAddress, udpReadError := connection.ReadFromUDP(buffer)

		if udpReadError == nil {
			go sendResponse(connection, remoteAddress, buffer)

		} else {
			log.Println(udpReadError)
			connection.SetWriteDeadline(time.Now().Add(time.Second * 1))
			connection.WriteToUDP([]byte("error "), remoteAddress)
		}
	}

}

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr, message []byte) {
	fmt.Println(string(message))
	conn.SetWriteDeadline(time.Now().Add(time.Second * 1))
	conn.WriteToUDP([]byte("Yes!!"), addr)
}
