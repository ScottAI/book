package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i:=0;i<20;i++{
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Print(" ",x)
		}(i)
	}
	fmt.Printf("\n%#v\n",wg)
	wg.Wait()
	fmt.Println("\nThe End!")
}
