package main

import (
	"flag"
	"log"
	"sync"
)

var messages = []string{
	"Hello!",
	"How are you?",
	"Are you just going to repeat what I say?",
	"So immature",
	"Stop copying me!",
}

// repeat concurrently prints out the given message n times
// parameter takes send only channel of strings

func repeat(n int, message string, c chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	// have a channel of strings
	// create n goroutines that send the message n times to the channel
	for i := 0; i < n; i++ {
		c <- message
	}
}

func main() {
	factor := flag.Int64("factor", 0, "The fan-out factor to repeat by")
	flag.Parse()

	if *factor <= 0 {
		log.Fatal("factor must be greater than 0")
	}

	// make waitgroup
	wg := sync.WaitGroup{}
	// add to waitgroup
	wg.Add(len(messages))
	// create a channel of strings
	ch := make(chan string, int(*factor)*len(messages))
	// create n goroutines that send the message n times to the channel
	for _, m := range messages {
		go repeat(int(*factor), m, ch, &wg)
		log.Println(<-ch)
	}
	// close the channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		log.Println(msg)
	}
}
