package main

import "fmt"

func main() {
	//方式一
	var funcList []func()
	for i:=0;i<3;i++{
		funcList = append(funcList,func(){
			fmt.Println(i)
		})
	}
	for _,f := range funcList{
		f()
	}
	//方式二
	var funcList2 []func()
	for i:=0;i<3;i++{
		j := i
		funcList2 = append(funcList2,func(){
			fmt.Println(j)
		})
	}
	for _,f := range funcList2{
		f()
	}
	//方式三
	var funcList3 []func(int)
	for i:=0;i<3;i++{
		funcList3 = append(funcList3,func(i int){
			fmt.Println(i)
		})
	}
	for i,f := range funcList3{
		f(i)
	}
}
