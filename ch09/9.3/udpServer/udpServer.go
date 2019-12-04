package main
import (
	"fmt"
	"net"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_,err := conn.WriteToUDP([]byte("From server: Hello, I got your mesage "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}


func main() {
	p := make([]byte, 512)
	addr := net.UDPAddr{
		Port: 7778,
		IP: net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	for {
		n,remoteaddr,err := ser.ReadFromUDP(p)
		fmt.Printf("Read a message from %v : %s \n", remoteaddr, p[0:n])
		if err !=  nil {
			fmt.Printf("Error:  %v", err)
			continue
		}
		go sendResponse(ser, remoteaddr)
	}
}