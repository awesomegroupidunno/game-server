package net

import (
	"fmt"
	"log"
	"net"
)

type UdpReceiver struct {
	PortNumber string
}

func (u *UdpReceiver) Run() {
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
		buffer := make([]byte, 2048)
		_, remoteAddress, udpReadError := connection.ReadFromUDP(buffer)

		if udpReadError == nil {
			go sendResponse(connection, remoteAddress, buffer)

		} else {
			connection.WriteToUDP([]byte("error "), remoteAddress)
		}
	}

}

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr, message []byte) {
	fmt.Println(string(message))
	conn.WriteToUDP([]byte("Yes!!"), addr)
}
