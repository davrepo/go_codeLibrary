package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := []int{2, 3, 4, 5, 6, 7, 8, 9}
	abort := make(chan bool)

	for _, num := range numbers {
		go calculateSquare(num, abort)
	}

	// Let's assume we want to stop calculations after 1 second
	time.Sleep(1 * time.Second)
	close(abort) // Broadcasting a signal to all goroutines to stop their work
}

func calculateSquare(num int, abort <-chan bool) {
	select {
	case <-abort:
		return // If we received a signal to stop, we just return from the function
	default:
		result := num * num
		fmt.Printf("Square of %d is %d\n", num, result)
		time.Sleep(2 * time.Second) // Simulating some delay
	}
}
