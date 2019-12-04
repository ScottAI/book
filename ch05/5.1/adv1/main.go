package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i:=10;i<20;i++{
			fmt.Print(" ",i)
		}
	}()
	fmt.Println()
	for i:=0;i<10;i++{
		fmt.Print(" ",i)
	}
	time.Sleep(2*time.Second)
}
