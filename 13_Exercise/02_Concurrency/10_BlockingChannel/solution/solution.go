package main

import (
	"flag"
	"log"
)

var messages = []string{
	"Hello!",
	"How are you?",
	"Are you just going to repeat what I say?",
	"So immature",
	"Stop copying me!",
}

// repeat concurrently prints out the given message n times
func repeat(n int, message string) {
	ch := make(chan struct{}) // signal channel
	for i := 0; i < n; i++ {
		go func(i int) {
			log.Printf("[G%d]:%s\n", i, message)
			ch <- struct{}{} // blocking until channel is read
		}(i)
	}

	// wait for all goroutines to finish
	for i := 0; i < n; i++ {
		<-ch // read from channel to unblock goroutine
	}

	// channel is a counting semaphore to wait for all goroutines to finish.
	// Each goroutine sends an empty struct to the channel when it completes,
	// and the 2nd loop waits until it has received n values from the channel,
	// which corresponds to all n goroutines completing.
}

func main() {
	factor := flag.Int64("factor", 0, "The fan-out factor to repeat by")
	flag.Parse()

	// repeat() itself blocks the main goroutine from processing the next message in the messages slice
	// until all repetitions of the current message have been printed.
	// This ensures that the messages are printed in the order they appear in the slice,
	// even though the individual repetitions of each message may be printed in any order
	// due to the concurrency of the goroutines.
	for _, m := range messages {
		log.Printf("[Main]:%s\n", m)
		repeat(int(*factor), m)
	}
}
