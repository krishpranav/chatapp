package main

import (
	"fmt"
	"net/http"

	"github.com/krishpranav/chatapp/web"
)

func serveWs(pool *web.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Websocket hit")
	conn, err := web.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &web.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func Routes() {
	pool := web.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

/* main func */
func main() {
	// Routes()
	fmt.Println("Chat App Started On: 8080")
	http.ListenAndServe(":8080", nil)
}
