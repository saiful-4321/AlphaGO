package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippets", snippets)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}

func home(w http.ResponseWriter, r *http.Request) {
	isAccessable := UrlRestriction(w, r, "/")

	if isAccessable {
		w.Write([]byte("Hello from home route"))
	}
}

func snippets(w http.ResponseWriter, r *http.Request) {
	isAccessable := UrlRestriction(w, r, "/snippets")

	if isAccessable {
		w.Write([]byte("Wellcome from snippets routes"))
	}
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	isAccessable := UrlRestriction(w, r, "/snippet/create")
	isMethodAllowed := CheckMethod(w, r, "POST")

	if isAccessable && isMethodAllowed {
		w.Write([]byte("Wellcome from create snippet route"))
	}
}

// url restriction process
func UrlRestriction(w http.ResponseWriter, r *http.Request, url string) bool {
	if r.RequestURI != url {
		http.NotFound(w, r)
		return false
	}
	return true
}

func CheckMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return false
	}
	return true
}
