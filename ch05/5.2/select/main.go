package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		//time.Sleep(3*time.Second)
		ch1 <- 1
	}()
	go func() {
		//time.Sleep(5*time.Second)
		ch2 <- 2
	}()

	select {
	case <-ch1:
		fmt.Println("can read from ch1")
	case <-ch2:
		fmt.Println("can read from ch2")
	//default:
	//	fmt.Println("default...")
	}
}
