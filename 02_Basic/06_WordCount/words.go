package main

import (
	"fmt"
	"strings"
)

func main() {
	// word count example
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	// fmt.Println(text)

	words := strings.Fields(text)
	counts := map[string]int{} // word -> count
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}

	fmt.Println(counts)
	// map[and:2 catch:1 me:2 needles:2 pins:2 sail:1 sew:1 the:1 to:1 wind:1]
}
