package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool) // channel to signal termination
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody) // only once
			done <- true      // flag termination
		}()
	}
	for i := 0; i < 10; i++ {
		<-done // wait for termination
	}
}
