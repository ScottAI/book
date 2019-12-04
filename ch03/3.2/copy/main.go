package main

import "fmt"

func main() {
	a1 := []int{1,1,1,1,1}
	b1 := []int{-1,-1,-1}
	copy(a1,b1)//将b1拷贝到a1
	fmt.Println("a1:",a1)
	fmt.Println("b1:",b1)

	a2 := []int{2,2,2,2,2}
	b2 := []int{-2,-2,-2}
	copy(b2,a2)//将a2拷贝到b2
	fmt.Println("a2:",a2)
	fmt.Println("b2",b2)
}
