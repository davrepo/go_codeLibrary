package main

import (
	"fmt"
	"time"
)

var hellos = []string{"Hello!", "Ciao!", "Hola!", "Hej!", "Salut!"}
var goodbyes = []string{"Goodbye!", "Arrivederci!", "Adios!", "Hej Hej!", "La revedere!"}

func main() {
	// create a channel
	ch := make(chan string, 1)
	ch2 := make(chan string, 1)
	// start the greeter to provide a greeting
	go greet(hellos, ch)    // ch 1
	go greet(goodbyes, ch2) // ch 2
	// sleep for a long time
	time.Sleep(1 * time.Second)
	fmt.Println("Main ready!")
	for {
		select {
		case gr, ok := <-ch:
			if !ok {
				ch = nil
			} else {
				printGreeting(gr)
			}
		case gr2, ok := <-ch2:
			if !ok {
				ch2 = nil
			} else {
				printGreeting(gr2)
			}
		default:
			if ch == nil && ch2 == nil {
				return
			}
		}
	}
}

// greet writes a greet to the given channel and then says goodbye
func greet(greetings []string, ch chan<- string) {
	fmt.Println("Greeter ready!")
	for _, g := range greetings {
		ch <- g
	}
	close(ch)
	fmt.Println("Greeter completed!")
}

// printGreeting sleeps and prints the greeting given
func printGreeting(greeting string) {
	// sleep and print
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Greeting received!", greeting)
}
