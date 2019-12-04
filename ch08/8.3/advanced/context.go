package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var (
	url string
	t = 5
	wg sync.WaitGroup
)

type information struct {
	r *http.Response
	err error
}

func connect(c context.Context) error  {
	defer wg.Done()
	info := make(chan information,1)
	tr := &http.Transport{}
	httpClient := &http.Client{Transport:tr}
	req,_ := http.NewRequest("GET",url,nil)
	req = req.WithContext(c)

	go func() {
		res,err := httpClient.Do(req)
		if err != nil {
			fmt.Println(err)
			info <- information{nil,err}
			return
		}else{
			info <- information{res,err}
		}
	}()

	select {
	case <- c.Done():
		fmt.Println("request is cancelled!!")
	case ok := <-info:
		err := ok.err
		r := ok.r
		if err != nil {
			fmt.Println("ERROR:",err)
			return err
		}
		defer r.Body.Close()
		realInfo,err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("ERROR:",err)
			return err
		}
		fmt.Printf("Response:%s\n",realInfo)
	}
	return nil
}

func main() {
	url = "http://google.com" //试着换成baidu
	ctx := context.Background()
	ctx,cancel := context.WithTimeout(ctx,time.Duration(t)*time.Second)
	defer cancel()
	fmt.Printf("connecting to %s \n",url)
	wg.Add(1)
	go connect(ctx)
	wg.Wait()
	fmt.Println("END...")
}
