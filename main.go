package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello from home route")
}
