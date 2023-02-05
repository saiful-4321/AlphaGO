package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello from home route"))
	})
	mux.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello from posts route"))
	})
	http.ListenAndServe("localhost:3000", mux)
}
