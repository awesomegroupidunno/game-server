package main
import (
	"fmt"
	"net"
	"bufio"
	"time"
)

func main() {

	for i:=0; i< 120000; i++ {
		p := make([]byte, 2048)
		conn, err := net.DialTimeout("udp", "172.31.1.42:10001", 1*time.Second)
		conn.SetDeadline(time.Now().Add(time.Second * 2))
		if err != nil {
			fmt.Printf("Some error %v", err)
			return
		}
		fmt.Fprintf(conn, "Hi %i", i)
		_, err = bufio.NewReader(conn).Read(p)
		if err == nil {
			fmt.Printf("%s, %i\n", p, i)
		} else {
			fmt.Printf("Some error %v\n", err)
		}
		conn.Close()
	}
}