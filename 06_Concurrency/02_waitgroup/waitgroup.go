package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go hello(&wg) // NB! pass pointer to wg to allow modification of wg, otherwise counter will not be decremented
	wg.Wait()
	goodbye()
}

func hello(wg *sync.WaitGroup) { // *sync.WaitGroup is a pointer to a WaitGroup
	defer wg.Done()
	fmt.Println("Hello")
}

func goodbye() {
	fmt.Println("Goodbye")
}
