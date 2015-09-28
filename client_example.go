package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
	"encoding/json"
	"github.com/awesomegroupidunno/game-server/cmd"
)

const host = "127.0.0.1:10001"
const seconds_timeout = 2
const message = "Hi "

func main() {

	//get connection to host and sets a timeout for connection when completed
	conn, err := net.DialTimeout("udp", host, 1*time.Second)
	for i := 0; i < 1200; i++ {
		response := make([]byte, 2048)

		if err != nil {
			log.Print(err)
		} else {
			goodmess, err := json.Marshal(&cmd.Command{Type:"GET", Subtype:"STATE", UniqueId:"ABC123"})
			//sends message to server
			//fmt.Fprintf(conn, goodmess, i)
			conn.Write(goodmess)


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
