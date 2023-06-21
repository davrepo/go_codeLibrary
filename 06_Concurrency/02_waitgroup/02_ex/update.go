package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func update(host, version string) {
	n := rand.Intn(100) + 50
	time.Sleep(time.Duration(n) * time.Millisecond)
	log.Printf("%s updated to %s", host, version)
}

func updateAll(version string, hosts <-chan string) { // receive-only channel
	var wg sync.WaitGroup
	for host := range hosts { // range over channel until it is closed
		wg.Add(1)
		go func(host, version string) {
			defer wg.Done()
			update(host, version)
		}(host, version)
	}

	wg.Wait()
}

func main() {
	ch := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			host := fmt.Sprintf("srv%d", i+1)
			ch <- host
		}
		// this is not necessary if we know loop will end at 5 iterations
		// but it's good practice to explicitly close the channel when done
		close(ch) // no more values will be sent on the channel
		// all receivers will be unblocked and receive zero values
	}()

	version := "1.0.2"
	updateAll(version, ch) // ch is a receive-only channel
	log.Printf("all servers updated")
}
