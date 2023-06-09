package main

import (
	"fmt"
	"time"
)

var greetings = []string{"Hello!", "Ciao!", "Hola!", "Hej!", "Salut!"}

func main() {
	ch := make(chan string, 1)
	go greet(ch)
	time.Sleep(5 * time.Second)
	fmt.Println("Main ready!") // (2)
	// greeting := <-ch will get blocked if channel in greet() is not closed
	for greeting := range ch {
		time.Sleep(2 * time.Second)
		fmt.Println("Greeting received!", greeting) // (3+/-)
	}
}

func greet(ch chan<- string) { // send only channel
	fmt.Println("Greeter ready!") // (1)
	for _, g := range greetings {
		ch <- g // send greeting to channel to use used by main()
	}
	close(ch)                         // close channel to signal no more data to main
	fmt.Println("Greeter completed!") // (3+/-)
}
