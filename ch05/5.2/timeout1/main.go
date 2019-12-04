package main

import (
	"fmt"
	"time"
)

func main() {
	//匿名函数中的代码块休眠4s
	ch1 := make(chan string)
	go func() {
		time.Sleep(4*time.Second)
		ch1 <- "ch1 si ready!"
	}()
	//注意time.After的使用，而且休眠时间为2s
	select {
	case mess := <- ch1:
		fmt.Println(mess)
	case t := <- time.After(2*time.Second):
		fmt.Println("ch1 timeout!",t)
	}
	//匿名函数休眠4s
	ch2 := make(chan string)
	go func() {
		time.Sleep(4*time.Second)
		ch2 <- "ch2 is ready!"
	}()
	//time.After内等待5s
	select {
	case mess := <- ch2:
		fmt.Println(mess)
	case t := <- time.After(5*time.Second):
		fmt.Println("ch2 timeout",t)
		
	}
}
