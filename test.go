package main

import "fmt"

func main() {
	x := 1
	n := &x         // &x is the address of x, assign to n, so n is a pointer to x
	fmt.Println(*n) // *n dereferences n, so *n is the value of x
	// 1
	*n = 50 // assign 50 to the value of n, which is dereferenced to x
	fmt.Println(x)
	// 50
}
