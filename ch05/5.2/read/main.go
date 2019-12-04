package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go writeChan(c,666)
	time.Sleep(1*time.Second)
	fmt.Println("Read:",<-c)
	if _,ok := <-c;ok{
		fmt.Println("Channel is Open")
	}else {
		fmt.Println("Channel is closed")
	}
}

func writeChan(c chan int,x int)  {
	fmt.Println("Start:",x)
	c <- x
	close(c)
	fmt.Println("End:",x)
}
