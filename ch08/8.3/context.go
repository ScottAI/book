package main

import (
	"context"
	"fmt"
	"time"
)

func testWCancel(t int)  {
	ctx := context.Background()
	ctx,cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		time.Sleep(3*time.Second)
		cancel()
	}()
	select {
		case <- ctx.Done():
			fmt.Println("testWCancel.Done:",ctx.Err())
		case e := <- time.After(time.Duration(t)*time.Second):
			fmt.Println("testWCancel:",e)
	}
	return
}

func testWDeadline(t int)  {
	ctx := context.Background()
	dl := time.Now().Add(time.Duration(1*t)*time.Second)
	ctx,cancel := context.WithDeadline(ctx,dl)
	defer cancel()
	go func() {
		time.Sleep(3*time.Second)
		cancel()
	}()
	select {
		case <- ctx.Done():
			fmt.Println("testWDeadline.Done:",ctx.Err())
		case e := <-time.After(time.Duration(t)*time.Second):
			fmt.Println("testWDeadline:",e)
	}
	return
}

func testWTimeout(t int)  {
	ctx := context.Background()
	ctx,cancel := context.WithTimeout(ctx,time.Duration(t)*time.Second)
	defer cancel()

	go func() {
		time.Sleep(3*time.Second)
		cancel()
	}()
	select {
		case <- ctx.Done():
			fmt.Println("testWTimeout.Done:",ctx.Err())
		case e := <-time.After(time.Duration(t)*time.Second):
			fmt.Println("testWTimeout:",e)
	}
	return
}

func main() {
	var t = 4
	testWCancel(t)
	testWDeadline(t)
	testWTimeout(t)
}
