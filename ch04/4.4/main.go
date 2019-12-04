package main

import (
	"fmt"
	"reflect"
)

type X struct {
	A1 int
	B1 float64
	C1 bool
}

type Y struct {
	A2 int
	B2 int
	C2 float64
	D2 string
}

func main() {
	x1 := X{A1:100,B1:3.14,C1:true}
	y1 := Y{A2:1,B2:2,C2:1.5,D2:"hello"}
	rx1 := reflect.ValueOf(&x1).Elem()
	ry1 := reflect.ValueOf(&y1).Elem()
	x1Type := rx1.Type()
	y1Type := ry1.Type()
	fmt.Printf("This type is %s,%d fileds of it are:\n",x1Type,rx1.NumField())
	for i:=0;i<rx1.NumField();i++{
		fmt.Printf("Name:%s,Type:%s,Value:%v\n",x1Type.Field(i).Name,rx1.Field(i).Type(),rx1.Field(i).Interface())
	}

	fmt.Printf("This type is %s,%d fields of it are:\n",y1Type,ry1.NumField())
	for i:=0;i<ry1.NumField();i++{
		fmt.Printf("Name:%s,Type:%s,Value:%v\n",y1Type.Field(i).Name,ry1.Field(i).Type(),ry1.Field(i).Interface())
	}
}
