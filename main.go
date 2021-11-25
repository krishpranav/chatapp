package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

/* upgrader */
// this is for read and write buffer size
var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

/* read messages to our websocket */
func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

/* define websocket */
func serve(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	reader(ws)
}

/* for now we can recive any response or request */
/* routes */
func Routes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})

	http.HandleFunc("/ws", serve)
}

/* main func */
func main() {
	Routes()
	fmt.Println("Chat App Started On: 8080")
	http.ListenAndServe(":8080", nil)
}