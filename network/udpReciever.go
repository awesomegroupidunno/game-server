package network

import (
	"github.com/awesomegroupidunno/game-server/encoder"
	"github.com/awesomegroupidunno/game-server/game"
	"github.com/awesomegroupidunno/game-server/state"
	"log"
	"net"
	"sync"
)

const default_max_packet = 8192
const default_port = ":10001"

var clientMutex sync.Mutex

type UdpReceiver struct {
	PortNumber     string
	MaxPacket      int
	EncoderDecoder encoder.EncoderDecoder
	connection     *net.UDPConn
	Router         game.CommandRouter
	Responses      chan state.StateResponse
	Acks           chan state.Ack
	clients        map[string]*net.UDPAddr
}

// Starts UDP Server
// Does Spawns 3 goroutines to deal with incoming and outgoing traffic
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
	clientMutex.Lock()
	u.clients = map[string]*net.UDPAddr{}
	clientMutex.Unlock()

	u.start()
}

// Actually starts UDP listening and brodcastring
// spawns 3 goroutines
func (u *UdpReceiver) start() {
	server_address, err_add := net.ResolveUDPAddr("udp", u.PortNumber)

	if err_add != nil {
		log.Fatal(err_add)
	}

	connection, errConn := net.ListenUDP("udp", server_address)

	if errConn != nil {
		log.Fatal(errConn)
	}
	u.connection = connection
	go u.startReceiver()
	go u.startSender()
	go u.startAcker()

}

// Runs in a loop to read incoming udp messages
func (u *UdpReceiver) startReceiver() {
	for {
		u.receiveUdp()
	}
}

// writes gamestate back to client
// consumes from u.Responses
func (u *UdpReceiver) startSender() {
	for {
		state := <-u.Responses
		clientMutex.Lock()
		for c := range u.clients {
			client := u.clients[c]

			for i := 0; i < len(state.State.Vehicles); i++ {
				if state.State.Vehicles[i].Owner == client.IP.String() {
					state.State.Vehicles[i].IsMe = true
				} else {
					state.State.Vehicles[i].IsMe = false
				}
			}

			buffer, error := u.EncoderDecoder.Encode(state.State)

			if error == nil {
				u.connection.WriteToUDP(buffer, client)
			} else {
				log.Println(error)
				u.connection.WriteToUDP([]byte("error"), client)
			}
		}
		clientMutex.Unlock()
	}
}

// writes acks back to client
// consumes from u.Acks
func (u *UdpReceiver) startAcker() {
	for {
		ack := <-u.Acks
		u.connection.WriteToUDP([]byte(ack.UUID), ack.Address)
	}
}

// listens for new udp packets
// decodes packet into commands
// forwards commands to router
func (u *UdpReceiver) receiveUdp() {
	buffer := make([]byte, u.MaxPacket)

	n, address, readError := u.connection.ReadFromUDP(buffer)
	a, encode_error := u.EncoderDecoder.Decode(buffer[:n])

	clientMutex.Lock()
	u.clients[address.String()] = address
	clientMutex.Unlock()

	if readError == nil && encode_error == nil {
		u.Router.RouteCommand(&a, address)
	} else {
		log.Println(readError)
		log.Println(encode_error)
		log.Println(address)
	}

}
