package main

import (
	"fmt"
	"net/http"
)

/* routes */
func Routes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
}

/* main func */
func main() {
	Routes()
	fmt.Println("Server started at 8080")
	http.ListenAndServe(":8080", nil)
}