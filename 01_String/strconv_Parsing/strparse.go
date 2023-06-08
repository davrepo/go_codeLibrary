package main

import (
	"fmt"
	"strconv"
)

func main() {
	sampleint := 100
	samplestr := "250"

	// This does a character conversion, not a numerical one
	newstr := string(sampleint)
	// The result is the character with the Unicode value of 100, which is d
	fmt.Println("Result of using string():", newstr) // Result of using string(): d

	// The strconv package contains a variety of functions for parsing and formatting
	// numbers, values, and strings

	// integer to string
	s := strconv.Itoa(sampleint)
	fmt.Printf("%T, %v\n", s, s) // string, 100

	// string to integer
	sampleint, err := strconv.Atoi(samplestr)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%T, %v\n", sampleint, sampleint) // int, 250

	// Other parse functions
	b, _ := strconv.ParseBool("true")
	fmt.Println(b)                            // true
	f, _ := strconv.ParseFloat("3.14159", 64) // float64
	fmt.Println(f)                            // 3.14159
	i, _ := strconv.ParseInt("-42", 10, 64)   // 10 is the base, 64 is the bit size, so int64
	fmt.Println(i)                            // -42
	u, _ := strconv.ParseUint("42", 10, 64)
	fmt.Println(u) // 42

	// given value => convert to string
	s = strconv.FormatBool(true)
	fmt.Println(s) // true
	s = strconv.FormatFloat(3.14159, 'E', -1, 64)
	fmt.Println(s) // 3.14159E+00
	s = strconv.FormatInt(-42, 10)
	fmt.Println(s)                 // -42
	s = strconv.FormatUint(42, 10) // 10 is the base, Uint is unsigned int
	fmt.Println(s)                 // 42
}
