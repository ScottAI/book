package main

import "fmt"

func main() {
	a := [...]int{1,2,3,4,5}
	ss := a[1:3]
	fmt.Println(a)
	fmt.Println(ss)
	ss[0] = 666
	ss[1] = -666
	fmt.Println(a)
	fmt.Println(ss)
	fmt.Println(cap(ss))
}
