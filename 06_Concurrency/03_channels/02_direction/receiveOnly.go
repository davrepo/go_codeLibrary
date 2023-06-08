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
	fmt.Println("Main ready!")
	for greeting := range ch {
		time.Sleep(2 * time.Second)
		fmt.Println("Greeting received!", greeting)
	}

}

// Send only channel
func greet(ch chan<- string) { // this is a send only channel
	fmt.Println("Greeter ready!")
	// greet
	for _, g := range greetings {
		ch <- g
	}
	close(ch)
	fmt.Println("Greeter completed!")
}
