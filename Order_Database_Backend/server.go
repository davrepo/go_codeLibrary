package main

// https://github.com/LinkedInLearning/applied-concurrency-in-go-3164282
// https://www.linkedin.com/learning/applied-concurrency-in-go/concurrency-in-daily-life?u=55937129

import (
	"fmt"
	"log"
	"net/http"

	"github.com/applied-concurrency-in-go/handlers"
)

func main() {
	fmt.Println("Welcome to the Orders App!")
	handler, err := handlers.New()
	if err != nil {
		log.Fatal(err)
	}
	// start server
	router := handlers.ConfigureHandler(handler)
	fmt.Println("Listening on localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
