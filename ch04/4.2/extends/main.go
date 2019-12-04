package main

import (
	"fmt"
	"image/color"
)

//定义一个矩形
type Rectangle struct{w,h float64}
//定义一个彩色矩形
type ColorRect struct {
	Rectangle
	Color color.RGBA
}
func main() {
	var cr ColorRect
	cr.h = 3
	cr.w = 2
	fmt.Println(cr.area())
}
//定义一个接收器为矩形类型的方法
func (r Rectangle) area() float64  {
	return r.w*r.h
}
