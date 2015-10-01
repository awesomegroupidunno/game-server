package network

import (
	"github.com/awesomegroupidunno/game-server/encoder"
	"github.com/awesomegroupidunno/game-server/game"
	"github.com/awesomegroupidunno/game-server/state"
	"log"
	"net"
)

const default_max_packet = 8192
const default_port = ":10001"

type UdpReceiver struct {
	PortNumber     string
	MaxPacket      int
	EncoderDecoder encoder.EncoderDecoder
	connection     *net.UDPConn
	Router         game.CommandRouter
	Responses      chan state.StateResponse
	Acks           chan state.Ack
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

	u.start()
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
	u.connection = connection
	go u.startReceiver()
	go u.startSender()
	go u.startAcker()

}

func (u *UdpReceiver) startReceiver() {
	for {
		u.receiveUdp()
	}
}

// writes gamestate back to client
func (u *UdpReceiver) startSender() {
	for {
		state := <-u.Responses
		buffer, error := u.EncoderDecoder.Encode(state.State)

		if error == nil {
			u.connection.WriteToUDP(buffer, state.Address)
		} else {
			log.Println(error)
			u.connection.WriteToUDP([]byte("error"), state.Address)
		}
	}
}

// writes acks back to client
func (u *UdpReceiver) startAcker() {
	for {
		ack := <-u.Acks
		u.connection.WriteToUDP([]byte(ack.Uuid), ack.Address)
	}
}

// listens for new udp packets
func (u *UdpReceiver) receiveUdp() {
	buffer := make([]byte, u.MaxPacket)

	n, address, readError := u.connection.ReadFromUDP(buffer)
	a, encode_error := u.EncoderDecoder.Decode(buffer[:n])

	if readError == nil && encode_error == nil {
		log.Println(a)
		u.Router.RouteCommand(&a, address)
	} else {
		log.Println(readError)
		log.Println(encode_error)
		log.Println(address)
	}

}
