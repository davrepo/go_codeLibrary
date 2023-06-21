package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is a GET")
}

func postRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is a POST")
}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is a DELETE")
}

func main() {
	r := mux.NewRouter()
	// curl http://localhost:8080/
	r.HandleFunc("/", getRequest).Methods("GET")
	// curl -X POST http://localhost:8080/
	r.HandleFunc("/", postRequest).Methods("POST")
	// curl -X DELETE http://localhost:8080/
	r.HandleFunc("/", deleteRequest).Methods("DELETE")

	http.Handle("/", r) // main endpoint, r is the router
	// if no method is specified, it defaults to GET
	fmt.Println("Server started and listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
