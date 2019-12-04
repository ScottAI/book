package main

import "fmt"

type data int


const (
	Zero data = iota
	One
	Two
	Three
	Four
)

func main()  {

	fmt.Println(Zero)
	fmt.Println(One)
	fmt.Println(Two)
	fmt.Println(Three)
	fmt.Println(Four)
}