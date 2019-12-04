package main

import "fmt"

func main() {
	for i:=0;i<10;i++{
		if i%2 == 0{
			continue
		}
		fmt.Print(i,"	")
	}
	fmt.Println()
	i := 5
	for{
		if i<1 {
			break
		}
		fmt.Print(i,"	")
		i--
	}

	fmt.Println()
	arr := []int{1,2,3,4,5}
	for i,v := range arr{
		fmt.Println("index:",i,"value:",v)
	}
}
