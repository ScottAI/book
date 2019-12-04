package main

import (
	"fmt"
	"math/rand"
)


func main() {
	//多个变量一起通过类型声明
	var i,j,k int
	fmt.Printf("i:%d,j:%d,k:%d",i,j,k)
	fmt.Println()
	//多个变量一起通过表达式声明
	var a,b,c = 1,"s",true
	fmt.Printf("a:%d,b:%s,c:%t",a,b,c)
	fmt.Println()
	//声明赋值的缩写
	f := rand.Float64()*100
	fmt.Printf("f:%g",f)
}
