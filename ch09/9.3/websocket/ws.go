package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// 1
type Message struct {
	Name string `json:"name"`
	Mess  string `json:"mess"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan *Message)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	// 2
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/chat", messHandler)
	http.HandleFunc("/ws", wsHandler)
	go echo()

	panic(http.ListenAndServe(":7778", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("/Users/liujinliang/projects/go/src/ljl/book/ch09/9.3/websocket/index.html")
	if err != nil {
		fmt.Println("Could not open file.", err)
	}
	fmt.Fprintf(w, "%s", content)
}

func writer(mess *Message) {
	broadcast <- mess
}

func messHandler(w http.ResponseWriter, r *http.Request) {
	var mess Message
	if err := json.NewDecoder(r.Body).Decode(&mess); err != nil {
		log.Printf("ERROR: %s", err)
		http.Error(w, "Bad request", http.StatusTeapot)
		return
	}
	defer r.Body.Close()
	go writer(&mess)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// register client
	clients[ws] = true
}

// 3
func echo() {
	for {
		mess := <-broadcast
		hisMess := fmt.Sprintf("%s : %s \n", mess.Name, mess.Mess)
		fmt.Println(hisMess)
		// send to every client that is currently connected
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(hisMess))
			if err != nil {
				log.Printf("Websocket error: %s", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}