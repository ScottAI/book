package main

import "fmt"

//定义一个矩形
type Rectangle struct{w,h float64}
func main() {
	rec := Rectangle{w:2,h:3}
	fmt.Println(area(rec.w,rec.h))
	fmt.Println(rec.area())
}

//定义一个矩形求面积的函数
func area(w,h float64)  float64{
	return w*h
}

//定义一个矩形类型的方法
func (r Rectangle) area() float64  {
	return r.w*r.h
}