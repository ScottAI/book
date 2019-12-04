package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)
type Logger struct { //Logger is a middleware handler
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	start := time.Now()
	l.handler.ServeHTTP(w,r)
	log.Printf("%s %s %v",r.Method,r.URL.Path,time.Since(start))
}

func NewLogger(handlerToWrap http.Handler) *Logger  {
	return &Logger{handlerToWrap}
}

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
	wrappedMux := NewLogger(mux) //包装了中间件的mux
	addr := "localhost:7774"
	log.Printf("listen at:%s",addr)
	log.Fatal(http.ListenAndServe(addr,wrappedMux))
}

