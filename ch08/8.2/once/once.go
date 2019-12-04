package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var once sync.Once
	for i:=0;i<10;i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(onlyOnce)
		}()
	}
	wg.Wait()
}

func onlyOnce()  {
	fmt.Println("only once")
}
