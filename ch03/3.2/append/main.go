package main

import "fmt"

func main() {
	var a = []int{1,2,3}
	fmt.Println("cap:",cap(a))
	a = append(a,333)
	fmt.Println("cap:",cap(a))
	a = append(a,[]int{-333,-333,-333,-333}...)
	fmt.Println("cap:",cap(a))

	for i,v := range a {
		fmt.Println(i,v)
	}

	//可以使用append进行删除
	a = append(a[:0],a[:3]...)//只保留前三个元素
	for i,v := range a {
		fmt.Println(i,v)
	}
	fmt.Println("cap:",cap(a))

	a = append([]int{222,222},a...)//在开头插入新的切片
	a = append(a[:2],append([]int{-222},a[2:]...)...)//在下标2的位置插入-222
	for i,v := range a {
		fmt.Println(i,v)
	}
	fmt.Println("cap:",cap(a))

}
