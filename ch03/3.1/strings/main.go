package main

import "fmt"

func main() {
	s := "hello,world"
	fmt.Println(len(s))
	fmt.Println(s[0],s[10])
	fmt.Println(s[:5])
	fmt.Println(s[6:])
	fmt.Println(s[:])
}
