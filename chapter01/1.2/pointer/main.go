package main

import "fmt"

func main() {
	m := 1
	selfPlusPointer(&m)
	fmt.Println(m)
	fmt.Println(*selfPlus(1))
}

func selfPlusPointer(n *int){
	*n++
}

func selfPlus(n int) *int {
	t := n+1
	return &t
}