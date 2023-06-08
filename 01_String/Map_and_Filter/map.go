package main

import (
	"fmt"
	"strings"
)

func main() {
	// The map function returns a copy of a string with the characters modified
	// according to the mapping function
	// this is a Caesar cipher example with a shift of 2
	shift := 2
	s := "The quick brown fox jumps over the lazy dog"

	// Create the mapping function
	transform := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z': // uppercase
			value := int('A') + (int(r) - int('A') + shift)
			if value > 91 { // 91 is the upper limit of uppercase characters
				value -= 26 // then wrap around
			} else if value < 65 {
				value += 26
			}
			return rune(value)
		case r >= 'a' && r <= 'z': // lowercase
			value := int('a') + (int(r) - int('a') + shift)
			if value > 122 {
				value -= 26
			} else if value < 97 {
				value += 26
			}
			return rune(value)
		}
		return r // if not an alphabet character, then return the same value
	}

	// Encode the message
	encode := strings.Map(transform, s)
	fmt.Println(encode)

	// Decode the message
	shift = -shift // call same function with negative shift
	decode := strings.Map(transform, encode)
	fmt.Println(decode)
}
