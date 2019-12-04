package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m1["k1"]=11
	m1["k2"]=22
	print(m1)
	delete(m1,"k1")
	print(m1)
	delete(m1,"k1")
	print(m1)

	val,ok := m1["k1"]
	if ok {
		fmt.Println(val)
	}else{
		fmt.Println("not exist")
	}

	var m2 map[string]int
	m2["kk1"] = -11
}

func print(m map[string]int)  {
	for k,v := range m{
		fmt.Println(k,v)
	}
}
