package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/awesomegroupidunno/game-server/cmd"
	"github.com/pborman/uuid"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

const host = "172.31.1.42:10001"
const seconds_timeout = 2

func main() {
	var waiter sync.WaitGroup
	waiter.Add(1)
	//get connection to host and sets a timeout for connection when completed
	conn, err := net.DialTimeout("udp", host, 1*time.Second)

	// start goroutine  to listen for messages
	go responsePrinter(conn)

	for i := 0; i < 12; i++ {

		if err != nil {
			log.Print(err)
		} else {
			var message []byte

			// creates unique identifier to be returned for the packet ack
			uuid := uuid.NewUUID().String()
			if i == 0 {
				a := cmd.NewConnect("Dan")
				a.UniqueId = uuid
				message, err = json.Marshal(a)
			} else {
				// randomly select between commands
				randType := rand.Intn(2)
				if randType == 0 {
					a := cmd.NewTurn(.5)
					a.UniqueId = uuid
					message, err = json.Marshal(a)
				}
				if randType == 1 {
					a := cmd.NewAcceleration(.5)
					a.UniqueId = uuid
					message, err = json.Marshal(a)
				}
			}

			//send the message
			conn.Write(message)

			time.Sleep(1 * time.Millisecond)
		}
	}
	waiter.Wait()

	//clean up connection
	conn.Close()

}

func responsePrinter(conn net.Conn) {
	for {
		// make a buffer for the response
		response := make([]byte, 2048)

		// read response
		_, err := bufio.NewReader(conn).Read(response)

		if err == nil {
			//print response
			fmt.Println(string(response))
		} else {
			log.Println(err)
		}
	}

}
