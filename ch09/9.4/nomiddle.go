package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func HelloHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Hello,World!"))
}

func CurrentTimeHandler(w http.ResponseWriter, r *http.Request)  {
	curTime := time.Now().Format(time.Kitchen)
	w.Write([]byte(fmt.Sprintf("Current time is %v",curTime)))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/hello",HelloHandler)
	mux.HandleFunc("/v1/time",CurrentTimeHandler)
	addr := "localhost:7774"
	log.Printf("listen at:%s",addr)
	log.Fatal(http.ListenAndServe(addr,mux))
}
