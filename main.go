package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		log.Fatal(err)
	}

	fmt.Printf("The id is %d", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	isAccessable := UrlRestriction(w, r, "/snippet/create")
	isMethodAllowed := CheckMethod(w, r, "POST")

	settingUpHeader(w, r)

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

// checking if the method is allowed or not
func CheckMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))
		http.Error(w, "Method not allowed", 405)
		return false
	}
	return true
}

// header customization while request
func settingUpHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ALLOW", "POST")
	w.Header().Set("custom", "test")
	w.Header().Set("Cache-Control", "public, max-age==31536000")
	w.Header().Add("Cache-Control", "public")
	w.Header().Add("Cache-Control", "max-age==31536000")
	// w.Header().Del("custom")
	// w.Header().Get("Cache-Control")
}
