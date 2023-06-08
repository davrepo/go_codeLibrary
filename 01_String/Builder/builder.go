package main

import (
	"fmt"
	"strings"
)

func main() {
	// Declare an empty strings.Builder
	var sb strings.Builder

	// Write some content
	sb.WriteString("This is string 1\n")
	sb.WriteString("This is string 2\n")
	sb.WriteString("This is string 3\n")

	// Output the concatenated string
	fmt.Println(sb.String())
	/*
	   This is string 1
	   This is string 2
	   This is string 3
	*/

	// Examine the builder's capacity
	fmt.Println("Capacity:", sb.Cap()) // Capacity: 96

	// Grow the capacity - use this if you know in advance how much data you need to write
	// you're going to be writing into the buffer to minimize copies
	sb.Grow(1024)
	// there is some overhead, so not exactly 1024
	fmt.Println("Capacity:", sb.Cap()) // Capacity: 1216

	for i := 0; i <= 10; i++ {
		// format a string and write it into sb, hence using &sb pointer
		fmt.Fprintf(&sb, "String %d -- ", i)
	}
	fmt.Println(sb.String())
	/*
		This is string 1
		This is string 2
		This is string 3
		String 0 -- String 1 -- String 2 -- String 3 -- String 4 -- String 5 -- String 6 -- String 7 -- String 8 -- String 9 -- String 10 --
	*/

	// we can get the length of what the final string will be
	fmt.Println("Builder size is", sb.Len()) // Builder size is 184

	// The Reset function will reset the builder to original state
	sb.Reset()
	fmt.Println("After Reset:")
	fmt.Println("Capacity:", sb.Cap())       // Capacity: 0
	fmt.Println("Builder size is", sb.Len()) // Builder size is 0
}
