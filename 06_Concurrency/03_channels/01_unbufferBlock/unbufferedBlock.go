package main

import (
	"fmt"
	"time"
)

// demo of unbuffered channel causing a block
func main() {
	// if you change below to buffered channel, it will not block
	// so ch := make(chan string, 1)
	ch := make(chan string)
	go greet(ch)
	time.Sleep(5 * time.Second)
	fmt.Println("Main ready!") // (2)
	greeting := <-ch
	time.Sleep(2 * time.Second)
	fmt.Println("Greeting received!") // (4)
	fmt.Println(greeting)             // (5)

}

// greet writes a greet to the given channel and then says goodbye
func greet(ch chan string) {
	fmt.Printf("Greeter ready!\nGreeter waiting to send greeting...\n") // (1)
	ch <- "Hello, world!"
	fmt.Println("Greeter completed!") // (3)
}
