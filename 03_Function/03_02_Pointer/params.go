package main

import (
	"fmt"
)

func main() {
	values := []int{1, 2, 3, 4} // slice of ints
	doubleAt(values, 2)
	fmt.Println(values) // [1 2 6 4]

	val := 10
	double(val)
	fmt.Println(val) // val is still 10

	doublePtr(&val)  // &val is the address of val
	fmt.Println(val) // val is now 20
}

func doubleAt(values []int, i int) {
	values[i] *= 2
	// this will change the value at the index
	// because slice is a pointer to an array
}

func double(n int) {
	n *= 2
	// this will not change the value of val
	// because n is a copy of val
}

func doublePtr(n *int) { // n is a pointer to an int
	*n *= 2 // dereference n and multiply value by 2
	// this will change the value at the address
}
