package main

import (
	"net/http"

	. "shopping/api"
)

func main() {
	// http.HandleFunc("/hello-world",world)
	// http.ListenAndServe(":8080",nil)

	srv := NewServer()
	http.ListenAndServe(":8080", srv)
}

func world(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
