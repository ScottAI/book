package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	//需要初始化随机数的资源库, 如果不执行这行, 不管运行多少次都返回同样的值
	rand.Seed(time.Now().UnixNano())
	no := rand.Intn(6)
	no *= 1000
	du := time.Duration(int32(no))*time.Millisecond
	fmt.Println("timeout duration is:",du)
	wg.Done()
	if isTimeout(&wg,du){
		fmt.Println("Time out!")
	}else {
		fmt.Println("Not time out")
	}
}

func isTimeout(wg *sync.WaitGroup,du time.Duration) bool  {
	ch1 := make(chan int)
	go func() {
		time.Sleep(3*time.Second)
		defer close(ch1)
		wg.Wait()
	}()
	select {
	case <-ch1:
		return false
	case <- time.After(du):
		return true
	}
}
