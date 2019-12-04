package main

import "fmt"

func main() {
	a := [...]int{1,2,3,4}
	b := &a
	fmt.Println(a[0],a[1])
	fmt.Println(b[0],b[1])
	var s = [...]string{"hello","世界"}
	for i,v := range s {
		fmt.Printf("s[%d]:%s\n",i,v)
	}
	for i:=0;i<len(s);i++{
		fmt.Printf("s[%d]:%s\n",i,s[i])
	}
}
