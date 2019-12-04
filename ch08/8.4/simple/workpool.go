package main

import (
	"fmt"
	"time"
)

func worcker(id int,jobs <- chan int,results chan <- int)  {
	for job := range jobs {
		fmt.Printf("worker(%d) start to do job(%d)\n",id,job)
		time.Sleep(time.Second)
		fmt.Printf("worker(%d) finished job(%d)\n",id,job)
		results <- job*job
	}
}

func main() {
	jobs := make(chan int,100)
	resultes := make(chan int,100)

	for id := 0;id<3;id++{
		go worcker(id,jobs,resultes)
	}
	for job:=1;job<=5;job++{
		jobs <- job
	}
	close(jobs)

	for i:=1;i<=5;i++{
		<-resultes
	}
}
