package main

import (
	"fmt"
	"time"
)

func main() {
	for i:=0;i<20;i++{
		go func(i int) {
			fmt.Print(" ",i)
		}(i)
	}
	time.Sleep(2*time.Second)
	fmt.Println("The End!")
}
