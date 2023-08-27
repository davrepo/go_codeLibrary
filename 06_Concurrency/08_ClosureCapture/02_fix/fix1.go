package main

import "fmt"

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func(val string) {
			fmt.Println(val)
			done <- true
		}(v)
	}

	// wait for all goroutines to complete before exiting
	for range values {
		<-done
	}
}
