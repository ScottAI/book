package main

import "fmt"

func main() {
	fmt.Println(swap(1,2))
	fmt.Println(sum(1,2,3,4,5,6,7,8,9,10))
	argus := []int{2,3,4,5,6,7,8,9,10}
	fmt.Println(sum(1,argus...))
	args := []interface{}{1234,"abcd"}
	print(args)
	print(args...)
}

func swap(a,b int) (int,int) {
	return b,a
}

func sum(a int,others ...int) int {
	for _,v := range others{
		a += v
	}
	return a
}

func print(a ...interface{})  {
	fmt.Println(a...)
}
