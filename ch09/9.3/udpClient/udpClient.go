package main
import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	p :=  make([]byte, 512)
	conn, err := net.Dial("udp", "127.0.0.1:7778")
	defer conn.Close()
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, "Hi UDP Server, How are you? \n")
	n, err := bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("Response:%s\n", p[0:n])
	} else {
		fmt.Printf("Error: %v\n", err)
	}
}