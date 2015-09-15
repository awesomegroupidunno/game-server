package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const host = "172.31.1.42:10001"
const seconds_timeout = 2
const message = "Hi "

func main() {

	for i := 0; i < 1200; i++ {
		response := make([]byte, 2048)

		//get connection to host and sets a timeout for connection when completed
		conn, err := net.DialTimeout("udp", host, 1*time.Second)
		conn.SetDeadline(time.Now().Add(time.Second * seconds_timeout))

		if err != nil {
			log.Print(err)
		} else {

			//sends message to server
			fmt.Fprintf(conn, message, i)

			_, err = bufio.NewReader(conn).Read(response)

			if err == nil {
				fmt.Println(string(response))
			} else {
				log.Println(err)
			}
		}

		conn.Close()
	}
}
