package main

import (
	"fmt"
	"sync"
)

var m sync.Map

func main() {
	//新增
	m.Store(1,"one")
	m.Store(2,"two")
	//LoadOrStore key不存在
	v,ok := m.LoadOrStore(3,"three")
	fmt.Println(v,ok)//three false
	//LoadOrStore key存在
	v,ok = m.LoadOrStore(1,"thisOne")
	fmt.Println(v,ok)//one true
	//Load
	v,ok = m.Load(1)
	if ok{
		fmt.Println("key is existed,and value is:",v)
	}else{
		fmt.Println("key is not existed")
	}
	//Range
	//先为Range准备一个传入的函数，该函数的声明格式是固定的，只有内部代
	//我们可以自定义
	f := func(k,v interface{}) bool{
		fmt.Println(k,v)
		return true
	}
	//执行Range
	m.Range(f)

	//Delete
	m.Delete(2)
	fmt.Println(m.Load(2)) //nil false
}
