package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// declare sample strings
	s := "the quick brown fox jumped over the lazy dog"
	s2 := []string{"one", "two", "three", "four"}
	s3 := "This is a string. With some punctionation, for a demo! Yep."

	// The Split function splits a string into substrings
	sub1 := strings.Split(s, " ")   // split s on spaces
	fmt.Printf("%q\n", sub1)        // ["the" "quick" "brown" "fox" "jumped" "over" "the" "lazy" "dog"]
	sub2 := strings.Split(s, "the") // split s on "the"
	fmt.Printf("%q\n", sub2)        // ["" " quick brown fox jumped over " " lazy dog"]

	// Join concatenates substrings, with the separator between each
	result := strings.Join(s2, " - ")
	fmt.Println(result) // one - two - three - four

	// Fields() is similar to Split() but splits on white space
	result2 := strings.Fields(s)
	// %q adds quotes around each string
	fmt.Printf("%q\n", result2) // ["the" "quick" "brown" "fox" "jumped" "over" "the" "lazy" "dog"]

	// FieldsFunc is a customizable version of fields that uses a callback
	// assign anonymous function to variable f
	// rune type is an alias for int32 and is used to distinguish character values from integer values
	f := func(c rune) bool {
		// unicode.IsPunct reports whether the rune is a Unicode punctuation character
		return unicode.IsPunct(c)
	}
	result3 := strings.FieldsFunc(s3, f)
	fmt.Printf("%q\n", result3) // ["This is a string" " With some punctionation" " for a demo" " Yep"]

	// Replacer can be used for multiple replacement operations
	rep := strings.NewReplacer(".", "|", ",", "|", "!", "|") // replace . , and ! with |
	result4 := rep.Replace(s3)
	fmt.Println(result4) // This is a string| With some punctionation| for a demo| Yep
}
