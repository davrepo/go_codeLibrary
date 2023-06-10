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
	// instead of using a waitgroup, we can use a channel of empty structs
	// this is a common pattern in Go
	// channel wait for all goroutines to finish before returning from the repeat function
	// which ensures all messages are printed before moving on to the next one
	// it isn't used to pass data between goroutines, hence the empty struct type
	// but rather to signal when a goroutine has completed
	ch := make(chan struct{})
	for i := 0; i < n; i++ {
		go func(i int) {
			log.Printf("[G%d]:%s\n", i, message)
			ch <- struct{}{} // ***
		}(i)
	}

	// wait for all goroutines to finish - this is blocking line
	// for loop reads from the channel n times. B/c channel reads are blocking in Go,
	// program will pause execution at this line if there is no data to read from the channel
	// loop will not complete until it has read n times from the channel,
	// which will only happen after all n goroutines have sent data to the channel
	for i := 0; i < n; i++ {
		<-ch // blocking
	}

	// channel is used as a counting semaphore to wait for all goroutines to finish.
	// Each goroutine sends an empty struct to the channel when it completes,
	// and the main goroutine waits until it has received n values from the channel,
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
