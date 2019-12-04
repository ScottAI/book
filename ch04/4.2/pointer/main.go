package main

import "fmt"

//定义一个矩形
type Rectangle struct{w,h float64}
func main() {
	p := &Rectangle{w:2,h:3}
	fmt.Println(p.area())

	rec := Rectangle{w:2,h:3}
	rp := &rec
	fmt.Println(rp.area())

	fmt.Println((&rec).area())
	fmt.Println(rec.area())    //会隐式的加上*rec，合法用法

	Rectangle{w:2,h:3}.area() //会报错，因为这种方式Go无法获取变量的地址
	Rectangle{w:2,h:3}.area2() //合法用法，因为area2方法的接收器不是指针类型，采用值传递


}
//定义一个接收器为指针类型的方法
func (r *Rectangle) area() float64  {
	return r.w*r.h
}
//定义一个接收器为矩形类型的方法
func (r Rectangle) area2() float64  {
	return r.w*r.h
}