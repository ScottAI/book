package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	ready = false
	singerNum = 3
)

func Sing(singerId int,c *sync.Cond)  {
	fmt.Printf("Singer (%d) is ready\n",singerId)
	c.L.Lock()
	for !ready {
		fmt.Printf("Singer (%d) is waiting\n",singerId)
		c.Wait()
	}
	fmt.Printf("Singer (%d) sing a song\n",singerId)
	ready = false
	c.L.Unlock()
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	for i:=0;i<singerNum;i++{
		go Sing(i,cond)
	}
	time.Sleep(3*time.Second)

	for i:=0;i<singerNum;i++{
		ready = true
		cond.Broadcast() //自行实验用Broadcast替换Signal方法的效果
		//cond.Signal()
		time.Sleep(3*time.Second)
	}
}
