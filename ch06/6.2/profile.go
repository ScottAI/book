package main

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var ch4cpu chan uint64
var chTimer chan struct{}
var memMap map[int]interface{}

func init(){
	ch4cpu = make(chan uint64, 10000)
	chTimer = make(chan struct{}, 20)
	memMap = make (map[int]interface{})
}
func main() {
	c, err := os.Create("/Users/liujinliang/projects/go/src/ljl/book/ch06/6.2/cpu_profile.prof")
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	m1, err := os.Create("/Users/liujinliang/projects/go/src/ljl/book/ch06/6.2/mem_profile1.prof")
	if err != nil {
		log.Fatal(err)
	}

	m2, err := os.Create("/Users/liujinliang/projects/go/src/ljl/book/ch06/6.2/mem_profile2.prof")
	if err != nil {
		log.Fatal(err)
	}

	m3, err := os.Create("/Users/liujinliang/projects/go/src/ljl/book/ch06/6.2/mem_profile3.prof")
	if err != nil {
		log.Fatal(err)
	}

	m4, err := os.Create("/Users/liujinliang/projects/go/src/ljl/book/ch06/6.2/mem_profile4.prof")
	if err != nil {
		log.Fatal(err)
	}

	defer m1.Close()
	defer m2.Close()
	defer m3.Close()
	defer m4.Close()

	pprof.StartCPUProfile(c)
	defer pprof.StopCPUProfile()

	memMap[1]= runMEMTest()

	runtime.GC()
	pprof.Lookup("heap").WriteTo(m1, 0)
	//从此处开始ch4cpu通道被不断的写入数据
	go runcputest()
	//goroutine运行15秒后chTimer写入值
	go func(){
		time.Sleep(15 * time.Second)
		log.Println("write timer")
		chTimer <- struct{}{}

	}()
	memMap[2]= runMEMTest()
	runtime.GC()
	pprof.Lookup("heap").WriteTo(m2, 0)

	memMap[2] = nil
	runtime.GC()
	pprof.Lookup("heap").WriteTo(m3, 0)

	memMap[1] = nil
	runtime.GC()
	pprof.Lookup("heap").WriteTo(m4, 0)

	procmsg()
}

func runMEMTest()([]int) {
	mem :=  make([]int, 100000, 120000)
	return mem
}

func runcputest(){
	var i uint64
	for {
		ch4cpu <- i
		i++
	}
}

func procmsg(){
	for {
		select {
		case _ = <-ch4cpu:
		case _ = <-chTimer://直到满足此条件for循环才结束
			log.Println("timeout")
			return
		}
	}
}