package main

import "fmt"

type Person struct {
	Name string
	Gender,Age int
}

type Employee struct {
	p Person
	Salary int
}

type Student struct {
	Person
	School string
}

func main() {
	e := Employee{p:Person{"Scott",1,30},Salary:1000}
	fmt.Println(e)
	fmt.Println(e.p.Name)

	var s Student
	s.Name = "Billy" //相当于 s.Person.Name = "Billy"
	s.Gender = 1	//相当于 s.Person.Gender = 1
	s.Age = 6		//相当于 s.Person.Age = 6
	s.School = "定慧里小学"
	fmt.Println(s)
}
