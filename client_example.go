package main
import (
	"fmt"
	"net"
	"bufio"
)

func main() {

	for i:=0; i< 120; i++ {
		p := make([]byte, 2048)
		conn, err := net.Dial("udp", "127.0.0.1:10001")
		if err != nil {
			fmt.Printf("Some error %v", err)
			return
		}
		fmt.Fprintf(conn, "Hi UDP Server, How are you doing? %i", i)
		_, err = bufio.NewReader(conn).Read(p)
		if err == nil {
			fmt.Printf("%s\n", p)
		} else {
			fmt.Printf("Some error %v\n", err)
		}
		conn.Close()
	}
}