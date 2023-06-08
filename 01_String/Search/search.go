package main

import (
	"fmt"
	"strings"
)

func main() {
	fname := "filename.txt"
	fname2 := "temp_picfile.jpeg"
	vowels := "aeiouAEIOU"
	s := "The quick brown fox jumps over the lazy dog"

	// Common string searching functions

	// is substring in a string
	fmt.Println(strings.Contains(s, "jump")) // true
	// if any of the given chars are in the string
	fmt.Println(strings.ContainsAny(s, "abcd")) // true

	// the offset of the first instance of a substring, -1 if not found
	fmt.Println(strings.Index(s, "fox")) // 16
	fmt.Println(strings.Index(s, "cat")) // -1
	// first instance of any of the given chars
	fmt.Println(strings.IndexAny("grzbl", vowels))   // -1
	fmt.Println(strings.IndexAny("Golang!", vowels)) // 1

	// HasPrefix and HasSuffix can be used to see if a string starts with
	// or ends with a specific substring
	fmt.Println(strings.HasSuffix(fname, "txt"))   // true
	fmt.Println(strings.HasPrefix(fname2, "temp")) // true

	// Count returns the number of non-overlapping instances of a substring
	fmt.Println(strings.Count(s, "the")) // 1
	// the with lower case the occur only once
	fmt.Println(strings.Count(s, "he")) // 2
	// he appears twice, The and the.
}
