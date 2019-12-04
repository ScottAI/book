package main

import "fmt"

type Person struct {
	Name string
	Gender,Age int
}

func main() {
	p1 := Person{Name:"Scott",Gender:1,Age:30}
	p2 := AddAge(p1)
	fmt.Println(p1)
	fmt.Println(p2)

	AddAgePlus(&p1) //注意参数
	fmt.Println(p1)

	pp := new(Person)
	AddAgePlus(pp)
	fmt.Println(pp)
}

func AddAge(p Person) (p2 Person){
	p.Age += 1
	return p
}
func AddAgePlus(pp *Person)  {
	pp.Age += 1
}

