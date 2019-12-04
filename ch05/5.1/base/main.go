package main

import (
	"fmt"
	"time"
)

func HelloWorld()  {
	fmt.Println("Hello,World!")
}

func main() {
	go HelloWorld()
	time.Sleep(1*time.Second)
	fmt.Println("The End!")
}
