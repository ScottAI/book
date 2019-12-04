package main

import (
	"fmt"
	"reflect"
)

type X struct {
	I int
	F float64
	S string
}
type Person struct {
	Name string `json:"jname"`
	Gender int	`json:"jgender"`
	Age int		`json:"jage"`
}

func (x X) CompareStr(xx X) bool {
	rx1 := reflect.ValueOf(&x).Elem()
	rx2 := reflect.ValueOf(&xx).Elem()
	for i:=0;i<rx1.NumField();i++{
		if rx1.Field(i).Interface() != rx2.Field(i).Interface(){
			return false
		}
	}
	return true
}

func (p Person) PrintTags()  {
	for i :=0;i<reflect.TypeOf(p).NumField();i++{
		fmt.Println(reflect.TypeOf(p).Field(i).Tag.Get("json"))
	}
}

func main() {
	x1 := X{I:1,F:1.2,S:"hello"}
	x2 := X{I:1,F:1.2,S:"hello"}
	fmt.Println(x1.CompareStr(x2))

	p := Person{Name:"Scott",Gender:1,Age:30}
	p.PrintTags()
	
}
