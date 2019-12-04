package main

import "net/http"

func main() {
	http.Handle("/", &ThisHandler{})
	http.Handle("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello！！"))
	}))
	http.HandleFunc("/hi", sayHi)

	http.ListenAndServe(":8080", nil)
}

func sayHi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi！！"))
}

type ThisHandler struct{}

func (m *ThisHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ThisHandler's ServeHTTP"))
}