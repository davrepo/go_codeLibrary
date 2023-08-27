package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	http.HandleFunc("/hello", Hello)
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
