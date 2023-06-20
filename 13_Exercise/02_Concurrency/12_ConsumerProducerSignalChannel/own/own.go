package main

import (
	"fmt"
	"log"
)

// given a list of recourses, simulates
// the concurrent allocation of recourses to consumer Goroutines.

// Signal channel: var abort chan struct{} = make(chan struct{})
// is a channel to synchronize Goroutines, not to send data.

// the number consumers
const consumerCount = 10

// types of resources to pass to consumers
var foodCourses = []string{
	"Caprese Salad",
	"Spaghetti Carbonara",
	"Vanilla Panna Cotta",
}

// Producer
func serveLunch(course string, out chan<- string, done <-chan struct{}) {
	// continuously send the course to the channel
	for { // infinite whilte true loop
		select {
		case out <- course: // send the course to the channel
		case <-done:
			// when done is closed,
			// it broadcasts a zero value and the case is selected
			return
		}
	}
}

// Consumer
func takeLunch(name string, in []chan string, done chan<- struct{}) {
	for _, ch := range in { // b/c in has 3 channels, so we loop 3 times
		log.Printf("%s eats %s.\n", name, <-ch)
	}
	done <- struct{}{} // abort that this consumer is done
}

// use slice of channels as a queue to hold the lunches
// which are produced and consumed concurrently
func main() {
	log.Printf("Welcome to cafeteria! Serving %d guests.\n", consumerCount)
	// slice of channels, total 3 channels, each channel is a course
	var courses []chan string
	doneEating := make(chan struct{})  // abort channel
	doneServing := make(chan struct{}) // abort channel
	for _, c := range foodCourses {
		ch := make(chan string)
		courses = append(courses, ch)     // append lunch to the slice of channels queue
		go serveLunch(c, ch, doneServing) // course string, single lunch channel, abort channel
		// since we're looping over the foodCourses slice,
		// we're creating goroutine for each particular course
	}
	for i := 0; i < consumerCount; i++ {
		name := fmt.Sprintf("Guest %d", i)
		// courses is a slice of channels, total 3 channels, each channel is a course
		go takeLunch(name, courses, doneEating) // name string, slice of channels, abort channel
	}

	for i := 0; i < consumerCount; i++ {
		<-doneEating // wait for all consumers to finish
	}
	close(doneServing) // close the abort channel to stop the producer
}
