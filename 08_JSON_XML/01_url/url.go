package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Define a URL
	s := "https://www.example.com:8000/user?username=joemarini"

	// Parse the URL content
	result, _ := url.Parse(s)    // Parse() returns a pointer to a URL struct
	fmt.Println(result.Scheme)   // https
	fmt.Println(result.Host)     // www.example.com:8000
	fmt.Println(result.Path)     // /user
	fmt.Println(result.Port())   // 8000
	fmt.Println(result.RawQuery) // username=joemarini

	// Extract the query components into a Values struct
	vals := result.Query()
	fmt.Println(vals)             // map[username:[joemarini]]
	fmt.Println(vals["username"]) // [joemarini]

	// create a URL from components
	newurl := &url.URL{
		Scheme:   "https",
		Host:     "www.example.com",
		Path:     "/args",
		RawQuery: "x=1&y=2",
	}
	// newurl is https://www.example.com/args?x=1&y=2
	s = newurl.String() // convert URL struct to a string
	fmt.Println(s)
	newurl.Host = "joemarini.com"
	s = newurl.String()
	fmt.Println(s)

	// Change the query parameters for the newurl URL
	// Values is a map[string][]string, is a type alias for map[string][]string
	newvals := url.Values{}
	newvals.Add("x", "100")
	newvals.Add("z", "somestr")
	newurl.RawQuery = newvals.Encode()
	// newurl is now https://www.example.com/args?x=100&z=somestr
	s = newurl.String()
	fmt.Println(s)
}
