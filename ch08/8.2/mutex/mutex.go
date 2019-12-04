package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m sync.Mutex
	v1 int
)

func Set(i int) {
	m.Lock()
	time.Sleep(time.Second)
	v1 =  i
	m.Unlock()
}
func Read() int {
	m.Lock()
	a := v1
	m.Unlock()
	return a

}
func main() {

	numGR := 5

	var wg sync.WaitGroup
	fmt.Printf("%d ", Read())
	for i := 1; i <= numGR; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Set(i)
			fmt.Printf("-> %d", Read())
		}(i)
	}

	wg.Wait()
}
