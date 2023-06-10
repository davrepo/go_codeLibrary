package main

import (
	"fmt"
	"log"
)

// given a list of recourses, simulates
// the concurrent allocation of recourses to consumer Goroutines.

// Signal channel: var signal chan struct{} = make(chan struct{})
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
		case out <- course:
		case <-done:
			// when done is closed,
			// it has a zero value of struct{} and the case is selected
			return
		}
	}
}

// Consumer
func takeLunch(name string, in []chan string, done chan<- struct{}) {
	for _, ch := range in {
		log.Printf("%s eats %s.\n", name, <-ch)
	}
	// this is same pattern as BlockingChannel exercise
	done <- struct{}{} // signal that this consumer is done
}

// use slice of channels as a queue to hold the lunches
// which are produced and consumed concurrently
func main() {
	log.Printf("Welcome to the conference lunch! Serving %d attendees.\n",
		consumerCount)
	var courses []chan string          // slice of channels to hold the lunches
	doneEating := make(chan struct{})  // signal channel
	doneServing := make(chan struct{}) // signal channel
	for _, c := range foodCourses {
		ch := make(chan string)
		courses = append(courses, ch)     // append lunch to the slice of channels queue
		go serveLunch(c, ch, doneServing) // current course string, single lunch channel, signal channel
		// one server per course
	}
	for i := 0; i < consumerCount; i++ {
		name := fmt.Sprintf("Attendee %d", i)
		go takeLunch(name, courses, doneEating) // name string, slice of channels, signal channel
	}

	for i := 0; i < consumerCount; i++ {
		<-doneEating // wait for all consumers to finish
	}
	close(doneServing) // close the signal channel to stop the producer
}
