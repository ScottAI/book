package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	addr := "www.baidu.com:80"
	tcpAddr,err := net.ResolveTCPAddr("tcp",addr)
	if err != nil {
		fmt.Println("Error:",err.Error())
		return
	}

	myConn,err := net.DialTCP("tcp",nil,tcpAddr)
	if err != nil {
		fmt.Println("Error:",err.Error())
		return
	}

	_,err = myConn.Write([]byte("HEAD / HTTP/1.1\r\n\r\n"))

	result,err := ioutil.ReadAll(myConn)
	if err != nil {
		fmt.Println("Error:",err.Error())
		return
	}
	fmt.Println(string(result))
	os.Exit(0)
}
