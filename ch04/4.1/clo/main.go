package main

import "fmt"

func main() {
	f := double()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func double() func() int  {
	var r int
	return func() int {
		r++
		return r*2
	}
}
