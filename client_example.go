package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/awesomegroupidunno/game-server/cmd"
	"log"
	"math/rand"
	"net"
	"time"
)

const host = "127.0.0.1:10001"
const seconds_timeout = 2

func main() {

	//get connection to host and sets a timeout for connection when completed
	conn, err := net.DialTimeout("udp", host, 1*time.Second)
	for i := 0; i < 1200; i++ {
		response := make([]byte, 2048)

		if err != nil {
			log.Print(err)
		} else {
			var message []byte

			randType := rand.Intn(3)
			if randType == 0 {
				a := cmd.NewTurn(.5)
				message, err = json.Marshal(a)
			}
			if randType == 1 {
				a := cmd.NewAcceleration(.5)
				message, err = json.Marshal(a)
			}
			if randType == 2 {
				a := cmd.NewState()
				message, err = json.Marshal(a)
			}

			conn.Write(message)

			_, err = bufio.NewReader(conn).Read(response)

			if err == nil {
				fmt.Println(string(response))
			} else {
				log.Println(err)
			}
		}
	}

	conn.Close()

}
