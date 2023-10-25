package main

import (
	"log"
	"strings"
	"time"
)

// strings.Split()
// strings.Join()
// GO string behaves like a slice of bytes []byte

const delay = 100 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	// split string into slice of words
	// let each word be a slice of bytes
	// each rune in the byte slice gets repeated by its index
	// and generates a new slice of bytes, which is the new word
	// join the words back into a string
	wordSlice := strings.Split(msg, " ")
	for _, word := range wordSlice {
		// split word into slice of bytes
		byteSlice := []byte(word)
		newByteSlice := make([]byte, len(byteSlice))
		for j, b := range byteSlice {
			count := j + 1
			// append byte to new byte slice j+1 times
			for k := 0; k < count; k++ {
				newByteSlice = append(newByteSlice, b)
			}
		}
		// append new word to new word slice
		print(string(newByteSlice))
	}
}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}
