package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippets", snippets)
	mux.HandleFunc("/snipet/create", createSnippet)

	log.Println("Starting server on :4000")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from home route"))
}

func snippets(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Wellcome from snippets routes"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Wellcome from create snippett route"))
}
