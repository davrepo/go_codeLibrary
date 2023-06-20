package main

import (
	"fmt"
)

func main() {
	// defer reads from bottom to top
	worker()
}

// output:
// worker
// Cleaning up B
// Cleaning up A

func worker() {
	r1, err := acquire("A")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer release(r1)

	r2, err := acquire("B")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer release(r2) // release B before A

	fmt.Println("worker")
}

func acquire(name string) (string, error) {
	return name, nil
}

func release(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}
