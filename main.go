package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Websocket hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

/* main func */
func main() {
	// Routes()
	fmt.Println("Chat App Started On: 8080")
	http.ListenAndServe(":8080", nil)
}
