package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("welcome to basic server section")

	http.ListenAndServe("localhost:3000", nil)
}
